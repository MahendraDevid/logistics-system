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
                git branch: 'main',
                    url: 'https://github.com/MahendraDevid/logistics-system.git'
            }
        }

        // ── WAREHOUSE SERVICE ──────────────────────────────────────

        stage('WMS - Unit Test') {
            steps {
                echo "=== Unit Test: Warehouse Service ==="
                dir('warehouse-service') {
                    // Menjamin semua library terdownload sebelum test
                    bat 'go mod tidy'
                    bat 'go test -v -count=1 ./internal/...'
                }
            }
        }

        stage('WMS - Lint') {
            steps {
                dir('warehouse-service') {
                    bat 'go vet ./...'
                }
            }
        }

        stage('WMS - Build Image') {
            steps {
                dir('warehouse-service') {
                    withCredentials([usernamePassword(
                        credentialsId: "${DOCKER_HUB_ID}",
                        passwordVariable: 'DOCKER_PASS',
                        usernameVariable: 'DOCKER_USER'
                    )]) {
                        // Username registry = DOCKER_HUB_USER (madeu30). Field Username di credential boleh salah; yang dipakai password/token saja.
                        powershell '$env:DOCKER_PASS | docker login -u $env:DOCKER_HUB_USER --password-stdin'
                        bat "docker build -t %DOCKER_HUB_USER%/warehouse-service:latest -f deployments/Dockerfile ."
                    }
                }
            }
        }

        stage('WMS - Functional Test') {
            steps {
                dir('warehouse-service') {
                    bat 'docker-compose -f deployments/docker-compose.test.yml up -d'
                    // Menunggu DB siap (15 detik)
                    bat 'timeout /t 15 /nobreak > nul'
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
                withCredentials([usernamePassword(
                    credentialsId: "${DOCKER_HUB_ID}",
                    passwordVariable: 'DOCKER_PASS',
                    usernameVariable: 'DOCKER_USER'
                )]) {
                    powershell '$env:DOCKER_PASS | docker login -u $env:DOCKER_HUB_USER --password-stdin'
                    bat "docker push %DOCKER_HUB_USER%/warehouse-service:latest"
                }
            }
        }

        // ── SETTLEMENT SERVICE ─────────────────────────────────────

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
                dir('settlement-service') {
                    bat 'go vet ./...'
                }
            }
        }

        stage('Settlement - Build Image') {
            steps {
                dir('settlement-service') {
                    withCredentials([usernamePassword(
                        credentialsId: "${DOCKER_HUB_ID}",
                        passwordVariable: 'DOCKER_PASS',
                        usernameVariable: 'DOCKER_USER'
                    )]) {
                        powershell '$env:DOCKER_PASS | docker login -u $env:DOCKER_HUB_USER --password-stdin'
                        bat "docker build -t %DOCKER_HUB_USER%/settlement-service:latest -f deployments/Dockerfile ."
                    }
                }
            }
        }

        stage('Settlement - Functional Test') {
            steps {
                dir('settlement-service') {
                    bat 'docker-compose -f deployments/docker-compose.test.yml up -d'
                    bat 'timeout /t 15 /nobreak > nul'
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
                withCredentials([usernamePassword(
                    credentialsId: "${DOCKER_HUB_ID}",
                    passwordVariable: 'DOCKER_PASS',
                    usernameVariable: 'DOCKER_USER'
                )]) {
                    powershell '$env:DOCKER_PASS | docker login -u $env:DOCKER_HUB_USER --password-stdin'
                    bat "docker push %DOCKER_HUB_USER%/settlement-service:latest"
                }
            }
        }

        // ── DEPLOYMENT ─────────────────────────────────────────────

        stage('Deploy to Kubernetes') {
            steps {
                withCredentials([file(credentialsId: "${KUBE_CONFIG_ID}", variable: 'KUBECONFIG_FILE')]) {
                    bat "kubectl --kubeconfig=%KUBECONFIG_FILE% apply -f warehouse-service/deployments/kubernetes/"
                    bat "kubectl --kubeconfig=%KUBECONFIG_FILE% apply -f settlement-service/deployments/kubernetes/"
                }
            }
        }

        stage('Verify Deployment') {
            steps {
                withCredentials([file(credentialsId: "${KUBE_CONFIG_ID}", variable: 'KUBECONFIG_FILE')]) {
                    bat "kubectl --kubeconfig=%KUBECONFIG_FILE% get pods"
                    bat "kubectl --kubeconfig=%KUBECONFIG_FILE% get svc"
                }
            }
        }
    }

    post {
        success {
            echo "✅ Pipeline berhasil! Kedua service sudah ter-deploy ke Kubernetes."
        }
        failure {
            echo "❌ Pipeline gagal! Periksa log build untuk detail error."
        }
    }
}