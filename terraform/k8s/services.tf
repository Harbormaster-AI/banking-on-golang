resource "kubernetes_service" "app-master" {
    metadata {
        name = "app-master"
    }

    spec {
        selector = {
          app  = "bankingOnGolang"
        }
        port {
            name        = "http"
            port        = 80
            target_port = ${}appPort}
        }

        port {
            name        = "db-port"
            port        = 3306
            target_port = 3306
        }

        port {
            port        = 8088
            target_port = 8088
            name        = "app-port"
        }

        type = "LoadBalancer"
    }
  
}
