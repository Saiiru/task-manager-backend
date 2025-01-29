pipeline {
    agent any

    environment {
        DOCKER_IMAGE = "your-docker-repo/your-backend-image"
        KUBECONFIG_CREDENTIALS = credentials('kconfig')
    }

    stages {
        stage('Checkout') {
            steps {
                git 'https://github.com/your-repo/your-backend.git'
            }
        }

        stage('Build') {
            steps {
                script {
                    docker.build(DOCKER_IMAGE)
                }
            }
        }

        stage('Test') {
            steps {
                sh 'go test ./... -cover'
            }
        }

        stage('Push') {
            steps {
                script {
                    docker.withRegistry('https://index.docker.io/v1/', 'dockerhub-credentials') {
                        docker.image(DOCKER_IMAGE).push()
                    }
                }
            }
        }

        stage('Deploy') {
            steps {
                script {
                    withKubeConfig([credentialsId: 'kubeconfig']) {
                        sh 'kubectl apply -f k8s/'
                    }
                }
            }
        }
    }

    post {
        always {
            cleanWs()
        }
    }
}
