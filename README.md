# backend
backend

remider how to commit :)

git commit -S -m "Ceitanya pesan"
git tag v1.0.0 -s -m "Ceritanya pesan"
git push origin main v1.0.0
go list -m github.com/FreightTrackr/backend@v1.0.0

docker build -t freight-trackr-backend .
gcloud init

Pick configuration to use:
 [1] Re-initialize this configuration [default] with new settings
 [2] Create a new configuration
Please enter your numeric choice:  1

Choose the account you would like to use to perform operations for this configuration:
 [1] befousmain@gmail.com
 [2] Log in with a new account
Please enter your numeric choice:  1

Pick cloud project to use: 
 [1] befous
 [2] plated-envoy-401509
 [3] ultra-concord-404810
 [4] Enter a project ID
 [5] Create a new project
Please enter numeric choice or text value (must exactly match list item):  1

net localgroup docker-users {nama komputer atau domain}\Befous /add
note: run as administrator, nama komputer bisa didapatkan di Control Panel > System and Security > System

gcloud auth login
gcloud auth configure-docker hostname (asia-southeast2-docker.pkg.dev)
docker tag freight-trackr-backend asia-southeast2-docker.pkg.dev/befous/gcf-artifacts/freight-trackr-backend
docker push asia-southeast2-docker.pkg.dev/befous/gcf-artifacts/freight-trackr-backend