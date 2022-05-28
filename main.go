package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"

	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

type Store struct {
	Context                    context.Context
	Namespaces                 []string
	CoreV1                     corev1.CoreV1Interface
	usedSecrets                []Item
	usedConfigMaps             []Item
	usedServiceAccounts        []Item
	usedPersistentVolumeClaims []Item
}

type Item struct {
	name       string
	namespaces string
}

func (k *Store) GetUsedResources() error {
	for _, v := range k.Namespaces {
		list, err := k.CoreV1.Pods(v).List(k.Context, metav1.ListOptions{Watch: false})
		if err != nil {
			return err
		}
		for _, vv := range list.Items {
			containers := vv.Spec.Containers
			for i := 0; i < len(containers); i++ {
				if containers[i].Env != nil {
					for ii := 0; ii < len(containers[i].Env); ii++ {
						if containers[i].Env[ii].ValueFrom != nil {
							if containers[i].Env[ii].ValueFrom.SecretKeyRef != nil {
								k.usedSecrets = append(k.usedSecrets, Item{containers[i].Env[ii].ValueFrom.SecretKeyRef.Name, vv.Namespace})
							}
						}
						if containers[i].Env[ii].ValueFrom != nil {
							if containers[i].Env[ii].ValueFrom.ConfigMapKeyRef != nil {
								k.usedConfigMaps = append(k.usedConfigMaps, Item{containers[i].Env[ii].ValueFrom.ConfigMapKeyRef.Name, vv.Namespace})
							}
						}
					}
				}
				if containers[i].EnvFrom != nil {
					for ii := 0; ii < len(containers[i].EnvFrom); ii++ {
						if containers[i].EnvFrom[ii].SecretRef != nil {
							k.usedSecrets = append(k.usedSecrets, Item{containers[i].EnvFrom[ii].SecretRef.Name, vv.Namespace})
						}
						if containers[i].EnvFrom[ii].ConfigMapRef != nil {
							k.usedConfigMaps = append(k.usedConfigMaps, Item{containers[i].EnvFrom[ii].ConfigMapRef.Name, vv.Namespace})
						}
					}
				}
			}
			volumes := vv.Spec.Volumes
			for i := 0; i < len(volumes); i++ {
				if volumes[i].Secret != nil {
					k.usedSecrets = append(k.usedSecrets, Item{volumes[i].Secret.SecretName, vv.Namespace})
				}
				if volumes[i].ConfigMap != nil {
					k.usedConfigMaps = append(k.usedConfigMaps, Item{volumes[i].ConfigMap.Name, vv.Namespace})
				}
				if volumes[i].PersistentVolumeClaim != nil {
					k.usedPersistentVolumeClaims = append(k.usedPersistentVolumeClaims, Item{volumes[i].PersistentVolumeClaim.ClaimName, vv.Namespace})
				}
			}
			serviceAccount := vv.Spec.ServiceAccountName
			if serviceAccount != "" {
				k.usedServiceAccounts = append(k.usedServiceAccounts, Item{serviceAccount, vv.Namespace})
			}
		}
	}
	return nil
}

func (k *Store) GetUnusedSecrets() ([]Item, error) {
	var secrets []Item
	for _, v := range k.Namespaces {
		list, err := k.CoreV1.Secrets(v).List(k.Context, metav1.ListOptions{Watch: false})
		if err != nil {
			return nil, err
		}
		for _, vv := range list.Items {
			secrets = append(secrets, Item{
				vv.Name,
				vv.Namespace,
			})
		}
	}
	return missingItems(k.usedSecrets, secrets), nil
}

func (k *Store) GetUnusedConfigMaps() ([]Item, error) {
	var configMaps []Item
	for _, v := range k.Namespaces {
		list, err := k.CoreV1.ConfigMaps(v).List(k.Context, metav1.ListOptions{Watch: false})
		if err != nil {
			return nil, err
		}
		for _, vv := range list.Items {
			configMaps = append(configMaps, Item{
				vv.Name,
				vv.Namespace,
			})
		}
	}
	return missingItems(k.usedConfigMaps, configMaps), nil
}

func (k *Store) GetUnusedServiceAccounts() ([]Item, error) {
	var serviceAccounts []Item
	for _, v := range k.Namespaces {
		list, err := k.CoreV1.ServiceAccounts(v).List(k.Context, metav1.ListOptions{Watch: false})
		if err != nil {
			return nil, err
		}
		for _, vv := range list.Items {
			serviceAccounts = append(serviceAccounts, Item{
				vv.Name,
				vv.Namespace,
			})
		}
	}
	return missingItems(k.usedServiceAccounts, serviceAccounts), nil
}

func (k *Store) GetUnusedPersistentVolumeClaims() ([]Item, error) {
	var persistentVolumeClaims []Item
	for _, v := range k.Namespaces {
		list, err := k.CoreV1.PersistentVolumeClaims(v).List(k.Context, metav1.ListOptions{Watch: false})
		if err != nil {
			return nil, err
		}
		for _, vv := range list.Items {
			persistentVolumeClaims = append(persistentVolumeClaims, Item{
				vv.Name,
				vv.Namespace,
			})
		}
	}
	return missingItems(k.usedPersistentVolumeClaims, persistentVolumeClaims), nil
}

type void struct{}

func missingItems(a, b []Item) []Item {
	ma := make(map[Item]void, len(a))
	var diffs []Item
	for _, ka := range a {
		ma[ka] = void{}
	}
	for _, kb := range b {
		if _, ok := ma[kb]; !ok {
			diffs = append(diffs, kb)
		}
	}
	return diffs
}

// Flags
type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var namespaces arrayFlags

func main() {
	// Flags
	flag.Var(&namespaces, "namespaces", "list of namespaces. (Required)")
	flag.Parse()
	// Test flags
	if len(namespaces) < 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	config := ctrl.GetConfigOrDie()
	clientSet := kubernetes.NewForConfigOrDie(config)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*200)
	defer cancel()
	var ns []string
	for _, v := range namespaces {
		ns = append(ns, v)
	}
	k := &Store{
		Context:    ctx,
		Namespaces: ns,
		CoreV1:     clientSet.CoreV1(),
	}
	err := k.GetUsedResources()
	if err != nil {
		log.Fatalln(err)
	}
	println("Unused Secrets")
	unusedSecrets, err := k.GetUnusedSecrets()
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range unusedSecrets {
		println("\t", v.name, v.namespaces)
	}
	println("Unused ConfigMaps")
	unusedConfigMaps, err := k.GetUnusedConfigMaps()
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range unusedConfigMaps {
		println("\t", v.name, v.namespaces)
	}
	println("Unused Service Accounts")
	unusedServiceAccounts, err := k.GetUnusedServiceAccounts()
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range unusedServiceAccounts {
		println("\t", v.name, v.namespaces)
	}

	println("Unused Persistent Volume Claims")
	unusedPersistentVolumeClaims, err := k.GetUnusedPersistentVolumeClaims()
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range unusedPersistentVolumeClaims {
		println("\t", v.name, v.namespaces)
	}
}
