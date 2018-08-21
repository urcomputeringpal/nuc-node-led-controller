/*
Copyright 2018 Jesse Newland

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	nodeutil "k8s.io/kubernetes/pkg/api/v1/node"
)

func setLed(key string) error {
	value := os.Getenv(fmt.Sprintf("NUC_LED_%s", strings.ToUpper(key)))
	bytes := []byte(value)
	log.Printf("Setting NUC LED to: %s", value)
	return ioutil.WriteFile("/proc/acpi/nuc_led", bytes, 0644)
}

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("greetings\n")

	for {
		node, err := clientset.CoreV1().Nodes().Get(os.Getenv("NODE_NAME"), metav1.GetOptions{})
		if err != nil {
			setLed("error")
		} else {
			if nodeutil.IsNodeReady(node) {
				if node.Spec.Unschedulable == true {
					setLed("unschedulable")
				} else {
					setLed("ready")
				}
			} else {
				setLed("not_ready")
			}
		}
		time.Sleep(10 * time.Second)
	}
}
