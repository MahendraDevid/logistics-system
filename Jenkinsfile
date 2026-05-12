pipeline {
    agent any

    environment {
        DOCKER_HUB_USER = 'madeu30'
        DOCKER_HUB_ID   = 'dockerhub-login'
        KUBE_CONFIG_ID  = 'kube-config'
    }

    stages {

        stage('Checkout Repository') {
            steps {
                echo "=== Mengambil kode terbaru ==="
                git branch: 'made',
                    url: 'https://github.com/MahendraDevid/logistics-system.git'
            }
        }

        stage('WMS - Unit Test') {
            steps {
                echo "=== Unit Test: Warehouse Service ==="
                dir('warehouse-service') {
                    bat 'go mod tidy'
                    bat 'go test -v -count=1 ./internal/...'
                }
            }
        }

        stage('WMS - Lint') {
            steps {
                echo "=== Lint: Warehouse Service ==="
                dir('warehouse-service') {
                    bat 'go vet ./...'
                }
            }
        }

        stage('WMS - Build Image') {
            steps {
                echo "=== Build Docker Image: Warehouse Service ==="
                dir('warehouse-service') {
                    withCredentials([usernamePassword(
                        credentialsId: "${DOCKER_HUB_ID}",
                        passwordVariable: 'DOCKER_PASS',
                        usernameVariable: 'DOCKER_USER'
                    )]) {
                        bat 'docker login -u %DOCKER_USER% -p %DOCKER_PASS%'
                        bat 'docker build -t %DOCKER_HUB_USER%/warehouse-service:latest -f deployments/Dockerfile .'
                    }
                }
            }
        }

        stage('WMS - Functional Test') {
            steps {
                echo "=== Functional Test: Warehouse Service ==="
                dir('warehouse-service') {
                    bat 'docker-compose -f deployments/docker-compose.test.yml up -d'
                    bat 'ping -n 16 127.0.0.1'
                    bat 'set TEST_DATABASE_URL=host=localhost user=testuser password=testpass dbname=wms_test port=5433 sslmode=disable && go test -v -tags=functional -count=1 ./tests/functional/...'
                }
            }
            post {
                always {
                    dir('warehouse-service') {
                        bat 'docker-compose -f deployments/docker-compose.test.yml down -v'
                    }
                }
            }
        }

        stage('WMS - Push Image') {
            steps {
                echo "=== Push Image: Warehouse Service ==="
                withCredentials([usernamePassword(
                    credentialsId: "${DOCKER_HUB_ID}",
                    passwordVariable: 'DOCKER_PASS',
                    usernameVariable: 'DOCKER_USER'
                )]) {
                    bat 'docker login -u %DOCKER_USER% -p %DOCKER_PASS%'
                    bat 'docker push %DOCKER_HUB_USER%/warehouse-service:latest'
                }
            }
        }

        stage('Settlement - Unit Test') {
            steps {
                echo "=== Unit Test: Settlement Service ==="
                dir('settlement-service') {
                    bat 'go mod tidy'
                    bat 'go test -v -count=1 ./internal/...'
                }
            }
        }

        stage('Settlement - Lint') {
            steps {
                echo "=== Lint: Settlement Service ==="
                dir('settlement-service') {
                    bat 'go vet ./...'
                }
            }
        }

        stage('Settlement - Build Image') {
            steps {
                echo "=== Build Docker Image: Settlement Service ==="
                dir('settlement-service') {
                    withCredentials([usernamePassword(
                        credentialsId: "${DOCKER_HUB_ID}",
                        passwordVariable: 'DOCKER_PASS',
                        usernameVariable: 'DOCKER_USER'
                    )]) {
                        bat 'docker login -u %DOCKER_USER% -p %DOCKER_PASS%'
                        bat 'docker build -t %DOCKER_HUB_USER%/settlement-service:latest -f deployments/Dockerfile .'
                    }
                }
            }
        }

        stage('Settlement - Functional Test') {
            steps {
                echo "=== Functional Test: Settlement Service ==="
                dir('settlement-service') {
                    bat 'docker-compose -f deployments/docker-compose.test.yml up -d'
                    bat 'ping -n 16 127.0.0.1'
                    bat 'set TEST_DATABASE_URL=host=localhost user=testuser password=testpass dbname=settlement_test port=5434 sslmode=disable && go test -v -tags=functional -count=1 ./tests/functional/...'
                }
            }
            post {
                always {
                    dir('settlement-service') {
                        bat 'docker-compose -f deployments/docker-compose.test.yml down -v'
                    }
                }
            }
        }

        stage('Settlement - Push Image') {
            steps {
                echo "=== Push Image: Settlement Service ==="
                withCredentials([usernamePassword(
                    credentialsId: "${DOCKER_HUB_ID}",
                    passwordVariable: 'DOCKER_PASS',
                    usernameVariable: 'DOCKER_USER'
                )]) {
                    bat 'docker login -u %DOCKER_USER% -p %DOCKER_PASS%'
                    bat 'docker push %DOCKER_HUB_USER%/settlement-service:latest'
                }
            }
        }

        stage('Deploy to Kubernetes') {
            steps {
                echo "=== Deploy ke Kubernetes ==="
                withCredentials([file(credentialsId: "${KUBE_CONFIG_ID}", variable: 'KUBECONFIG_FILE')]) {
                    bat 'kubectl --kubeconfig=%KUBECONFIG_FILE% apply -f warehouse-service/deployments/kubernetes/'
                    bat 'kubectl --kubeconfig=%KUBECONFIG_FILE% apply -f settlement-service/deployments/kubernetes/'
                }
            }
        }

        stage('Verify Deployment') {
            steps {
                echo "=== Verifikasi Deployment ==="
                withCredentials([file(credentialsId: "${KUBE_CONFIG_ID}", variable: 'KUBECONFIG_FILE')]) {
                    bat 'kubectl --kubeconfig=%KUBECONFIG_FILE% get pods'
                    bat 'kubectl --kubeconfig=%KUBECONFIG_FILE% get svc'
                }
            }
        }
    }

    post {
        success {
            echo "Pipeline berhasil! Kedua service sudah ter-deploy ke Kubernetes."
        }
        failure {
            echo "Pipeline gagal! Periksa log build untuk detail error."
        }
    }
}