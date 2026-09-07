resource "kubernetes_service" "app-master" {
    metadata {
        name = "app-master"
    }

    spec {
        selector = {
          app  = "bankingongolang"
        }
        port {
            name        = "http"
            port        = 80
            target_port = 4000
        }

        port {
            name        = "db-port"
            port        = 
            target_port = 
        }

        port {
            port        = 4000
            target_port = 4000
            name        = "app-port"
        }

        type = ""
    }
  
}
