// package api

// import (
// 	"context"
// 	"encoding/json"
// 	"flag"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/gorilla/mux"
// 	"github.com/joho/godotenv"
// 	_ "github.com/lib/pq"
// 	db "github.com/mrudraia/k8s-postgres-containerised/db"
// 	models "github.com/mrudraia/k8s-postgres-containerised/models"
// 	"gorm.io/gorm"
// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// 	"k8s.io/client-go/kubernetes"
// 	"k8s.io/client-go/rest"
// )

// type Repository struct {
// 	DB *gorm.DB
// }

// type PodsStruct struct {
// 	Name      string
// 	Count     int
// 	Namespace string
// 	Id        string
// }

// func (r *Repository) Run() error {
// 	var ns string
// 	flag.StringVar(&ns, "namespace", "", "namespace")
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	config := &db.Config{
// 		Host:     os.Getenv("DB_HOST"),
// 		Port:     os.Getenv("DB_PORT"),
// 		Password: os.Getenv("DB_PASS"),
// 		User:     os.Getenv("DB_USER"),
// 		SSLMode:  os.Getenv("DB_SSLMODE"),
// 		DBName:   os.Getenv("DB_NAME"),
// 	}

// 	db, err := db.NewConnection(config)

// 	if err != nil {
// 		log.Fatal("Could not load the database")
// 	}
// 	err1 := models.MigratePods(db)
// 	err2 := models.MigrateServices(db)

// 	if err1 != nil || err2 != nil {
// 		log.Fatal("could not migrate db")
// 	}

// 	podModel := &[]models.Pods{}
// 	pod_table := r.DB.Find(podModel).Error
// 	if err != nil {
// 		// context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get books"})
// 		return err
// 	} else {
// 		fmt.Println(pod_table)
// 	}

// 	serviceModel := &[]models.Pods{}
// 	service_table := r.DB.Find(serviceModel).Error
// 	if err != nil {
// 		// context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get books"})
// 		return err
// 	} else {
// 		fmt.Println(service_table)
// 	}

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	kube_config, err := rest.InClusterConfig()
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	clientset, err := kubernetes.NewForConfig(kube_config)
// 	if err != nil {
// 		fmt.Printf("error %s, creating clientset\n", err.Error())
// 	}

// 	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})

// 	if err != nil {
// 		fmt.Printf("error %s, while listing all the pods from default namespace\n", err.Error())
// 	}
// 	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

// 	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		fmt.Fprintf(w, "hello world testing")
// 		return
// 	})

// 	http.HandleFunc("/pods", func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		mySlice := []PodsStruct{}
// 		for i, pod := range pods.Items {
// 			mySlice = append(mySlice, PodsStruct{Name: pod.GetName(), Count: i, Namespace: pod.GetNamespace(), Id: string(pod.UID)})
// 		}
// 		b, err := json.MarshalIndent(mySlice, "", "    ")
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		w.Write(b)

// 		return
// 	})

// 	http.HandleFunc("/pods/{id}", func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		params := mux.Vars(r)
// 		for _, pod := range pods.Items {
// 			if params["id"] == pod.Name {
// 				b, err := json.Marshal(PodsStruct{Name: pod.GetName(), Namespace: pod.GetNamespace()})
// 				if err != nil {
// 					panic(err)
// 				}
// 				w.Write(b)
// 			}
// 		}

// 		return
// 	})

// 	http.ListenAndServe(":9090", nil)

// 	return nil

// }
