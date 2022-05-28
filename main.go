package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"

	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	"github.com/jessequinn/k8s-resource-finder/internal"
)

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
	k := internal.New(ctx, ns, clientSet.CoreV1())
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
		println("\t", v.Name, v.Namespaces)
	}
	println("Unused ConfigMaps")
	unusedConfigMaps, err := k.GetUnusedConfigMaps()
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range unusedConfigMaps {
		println("\t", v.Name, v.Namespaces)
	}
	println("Unused Service Accounts")
	unusedServiceAccounts, err := k.GetUnusedServiceAccounts()
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range unusedServiceAccounts {
		println("\t", v.Name, v.Namespaces)
	}

	println("Unused Persistent Volume Claims")
	unusedPersistentVolumeClaims, err := k.GetUnusedPersistentVolumeClaims()
	if err != nil {
		log.Fatalln(err)
	}
	for _, v := range unusedPersistentVolumeClaims {
		println("\t", v.Name, v.Namespaces)
	}
}
