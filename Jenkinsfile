pipeline{
    environment {
        registry = "nimrodops/godockerapp"
        registryCredential = 'docker-credentials'
        dockerImage = ''
    }
    agent any

    stages{
        stage("build image"){
            steps{
                echo "========executing build image========"
                script{
                    dockerImage = docker.build registry + ":$BUILD_NUMBER"
                }
            }
        }

        stage("deploy image to docker hub"){

            steps{
                echo "========executing deploy image to docker hub========"
                script{
                    docker.withRegistry('', registryCredential) {
                        dockerImage.push()
                    }
                }
            }
        }

        stage('Remove Unused docker image'){

            steps{
                echo "========executing Remove unused docker image========"
                sh "docker rmi $registry:$BUILD_NUMBER"
            }
        }
    }


}
