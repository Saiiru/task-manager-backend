pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                withCredentials([
                    string(credentialsId: 'dockerhub', variable: 'DOCKER_USERNAME'),
                    string(credentialsId: 'dockerhub_password', variable: 'DOCKER_PASSWORD'),
                    string(credentialsId: 'db_user', variable: 'DB_USER'),
                    string(credentialsId: 'db_password', variable: 'DB_PASSWORD'),
                    string(credentialsId: 'db_name', variable: 'DB_NAME'),
                    string(credentialsId: 'jwt_secret', variable: 'JWT_SECRET')
                ]) {
                    sh 'docker build -t $DOCKER_USERNAME/task-manager-backend .'
                }
            }
        }
        stage('Push') {
            steps {
                withCredentials([
                    string(credentialsId: 'dockerhub', variable: 'DOCKER_USERNAME'),
                    string(credentialsId: 'dockerhub_password', variable: 'DOCKER_PASSWORD')
                ]) {
                    sh 'docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD'
                    sh 'docker push $DOCKER_USERNAME/task-manager-backend'
                }
            }
        }
        stage('Deploy') {
            steps {
                withCredentials([file(credentialsId: 'kubeconfig', variable: 'KUBECONFIG')]) {
                    sh 'kubectl apply -f k8s/'
                }
            }
        }
    }
}