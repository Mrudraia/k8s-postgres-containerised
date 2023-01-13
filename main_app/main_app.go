package main_app

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	db "github.com/mrudraia/k8s-postgres-containerised/db"
	"github.com/mrudraia/k8s-postgres-containerised/models"
	"gorm.io/gorm"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Repository struct {
	DB *gorm.DB
}

func Application() {
	var ns string
	flag.StringVar(&ns, "namespace", "", "namespace")

	config := &db.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	sqlDB, _ := db.NewConnection(config)
	err1 := models.MigratePods(sqlDB)
	err2 := models.MigrateServices(sqlDB)
	if err1 != nil || err2 != nil {
		log.Fatal("could not migrate db")
	}

	configration, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(configration)
	if err != nil {
		log.Fatal(err)
	}

	pods, err := clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})

	if err != nil {
		log.Fatalln("Failed to get pods:", err)
	}

	services, err := clientset.CoreV1().Services("").List(context.Background(), metav1.ListOptions{})

	if err != nil {
		log.Fatalln("Failed to get services:", err)
	}

	deployments, err := clientset.AppsV1().Deployments("").List(context.Background(), metav1.ListOptions{})

	if err != nil {
		log.Fatalln("Failed to get deployments:", err)
	}

	http.HandleFunc("/pods", func(w http.ResponseWriter, r *http.Request) {
		for i, pod := range pods.Items {
			cont := models.Pods{Name: pod.GetName(), Namespace: pod.GetNamespace(), Count: strconv.Itoa(i)}
			sqlDB.Create(&cont)
			fmt.Fprintf(w, "[%d] %s\n", i, pod.GetName())
		}
	})

	http.HandleFunc("/services", func(w http.ResponseWriter, r *http.Request) {
		for i, ser := range services.Items {
			serv := models.Services{Name: ser.GetName(), Type: string(ser.Spec.Type), Count: strconv.Itoa(i)}
			sqlDB.Create(&serv)
			fmt.Fprintf(w, "[%d] %s\n", i, ser.GetName())
		}
	})

	http.HandleFunc("/deployments", func(w http.ResponseWriter, r *http.Request) {
		for i, dep := range deployments.Items {
			fmt.Fprintf(w, "[%d] %s\n", i, dep.GetName())
		}
	})

	http.ListenAndServe(":9090", nil)

}
