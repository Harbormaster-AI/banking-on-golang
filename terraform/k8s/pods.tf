resource "kubernetes_replication_controller" "app-master" {
    metadata {
        name = "app-master"
    }

    spec {
        replicas = 1

        selector = {
            app  = "bankingOnGolang"
        }

#Declare_K8_Containers()

    }
}