pipeline {
    agent any
    environment {
        KUBECONFIG = credentials('kubeconfig')
        DOCKER_USERNAME = credentials('dockerhub').username
        DOCKER_PASSWORD = credentials('dockerhub').password
        DB_USER = credentials('db_user')
        DB_PASSWORD = credentials('db_password')
        DB_NAME = credentials('db_name')
        JWT_SECRET = credentials('jwt_secret')
    }
    stages {
        stage('Build') {
            steps {
                sh 'docker build -t $DOCKER_USERNAME/task-manager-backend .'
            }
        }
        stage('Push') {
            steps {
                sh 'docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD'
                sh 'docker push $DOCKER_USERNAME/task-manager-backend'
            }
        }
        stage('Deploy') {
            steps {
                sh 'kubectl apply -f k8s/'
            }
        }
    }
}