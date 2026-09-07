
# Modules
module "eks" {
  source   = "./eks"
}

module "k8s" {
  source   = "./k8s"
}


provider "kubernetes" {
    host = $k8Host
    cluster_ca_certificate = base64decode(
     aws_eks_cluster.this.certificate_authority[0].data
    )

    exec {
      api_version = "client.authentication.k8s.io/v1"

      command = "aws"

      args = [
        "eks",
        "get-token",
        "--cluster-name",
        aws_eks_cluster.this.name
      ]
    }
}