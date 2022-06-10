### Herramientas usadas
* Terraform (v0.12.7)
* Shell 
* Helm (v2.14.1)
* Kubernetes (1.13)

### Servicios de AWS usados 
* EKS (Elastic Kubernetes Service)
* ECR (Elastic Container Registry)


### Descripción del Pipeline de despliegue

El pipeline desarrollado es muy sencillo.

Se consideró solo el ambiente de DEV. 

##### Pipeline Steps
* Se activa la compilación con Jenkins y se siguen los siguientes pasos: 
	* Build: se crea una imagen acoplable con una etiqueta, esta etiqueta será jenkins BUILD_NUMBER.
	* Push Image: Se hace push de esta imagen al repositorio de ECR.
	* Integration Test: en esta prueba, implementamos la imagen en un espacio de nombres diferente y hacemos las pruebas en ese espacio de nombres.
	* Si la prueba es exitosa, se implementará para liberar el espacio de nombres; de lo contrario, se cancelará la implementación.  


### Instrucciones

##### Prerequisitos 
* Docker 
* Kubernetes 
* awscli 
* helm
* terraform

##### Despliegue de la infraestructura 
Ubicados en el directorio de la aplicación, se ejecutan los siguientes somandos:  

``` terraform plan ```

luego se ejecuta

``` terraform apply ```


##### Definiendo el ambiente - clusters 

1. Definición del config de kubernetes
   `aws eks --region eu-central-1 update-kubeconfig --name eks-k8s-demo-dev-cluster`
   
2. Inicializar helm 
	
	`helm init `

3. Login al ECR 

	`$(aws ecr get-login --no-include-email --region eu-central-1)`
		
4. Verificación de los nodos 

	`kubectl get nodes`

5. Definición de las variables de ambiente sobre el ECR 
   
   `export REPOSITORY_URL="<replace with ecr repo url>"`
   

##### Primero 
Verificar las variables de ambiente 

`echo $REPOSITORY_URL`

Ejecutar el script 

`sh ci/firsttime.sh`


##### Build Script 
Ejecución del pipeline de jenkins

`sh ci/build.sh <BUILD_NUMBER>`

`BUILD_NUMBER` jenkins build number.
