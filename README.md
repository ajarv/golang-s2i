
# An OpenShift Source-to-Image builder dedicated to Go  

## Installation / usage

### Using Docker

#### Create the builder image
The following command will create a builder image named golang-centos7.
```
docker build -t golang-centos7 .
```

#### Testing the builder image
The builder image can be tested using the following commands:
```
docker build -t golang-centos7-candidate .
IMAGE_NAME=golang-centos7-candidate test/run
```

#### Creating the application image
The application image combines the builder image with your applications source code, which is served using whatever application is installed via the *Dockerfile*, compiled using the *assemble* script, and run using the *run* script.
The following command will create the application image:
```
s2i build test/test-app golang-centos7 golang-centos7-app
```
Using the logic defined in the *assemble* script, s2i will now create an application image using the builder image as a base and including the source code from the test/test-app directory.

#### Running the application image
Running the application image is as simple as invoking the docker run command:
```
docker run -d -p 8080:8080 golang-centos7-app
```
The application, which consists of a simple static web page, should now be accessible at  [http://localhost:8080](http://localhost:8080).

### Using OpenShift

#### Create the builder image
The following command will create a builder image named golang-centos7.

```
$ oc new-build https://github.com/nmasse-itix/golang-s2i.git --name=golang-centos7
$ oc logs -f bc/golang-centos7
```

OR

```
$ oc login -u admin..
$ oc project openshift
$ oc new-build https://github.com/nmasse-itix/golang-s2i.git --name=golang-centos7 --build-arg=GOLANG_VERSION=1.10.4
$ oc logs -f bc/golang-centos7
```


#### Creating and running the application image
The application image combines the builder image with your applications source code, which is served using whatever application is installed via the *Dockerfile*, compiled using the *assemble* script, and run using the *run* script.
The following command will create the application image and run it:
```
$ oc new-app --strategy=source --context-dir=test/test-app --name golang-centos7-app golang-centos7~https://github.com/nmasse-itix/golang-s2i.git
$ oc logs -f bc/golang-centos7-app
```
Using the logic defined in the *assemble* script, s2i will now create an application image using the builder image as a base and including the source code from the test/test-app directory.

Creating a route for the application is needed to use it :
```
$ oc expose service golang-centos7-app
$ oc get route golang-centos7-app
```

The application, which consists of a simple static web page, should now be accessible at the url shown on last command (oc get route) :
```
$ curl "http://$(oc get route golang-centos7-app |awk '$1 == "golang-centos7-app" {print $2}')/"
```



## Files and Directories  
| File                   | Required? | Description                                                  |
|------------------------|-----------|--------------------------------------------------------------|
| Dockerfile             | Yes       | Defines the base builder image                               |
| s2i/bin/assemble       | Yes       | Script that builds the application                           |
| s2i/bin/usage          | No        | Script that prints the usage of the builder                  |
| s2i/bin/run            | Yes       | Script that runs the application                             |
| s2i/bin/save-artifacts | No        | Script for incremental builds that saves the built artifacts |
| test/run               | No        | Test script for the builder image                            |
| test/test-app          | Yes       | Test application source code                                 |

### Dockerfile
Create a *Dockerfile* that installs all of the necessary tools and libraries that are needed to build and run our application.  This file will also handle copying the s2i scripts into the created image.

### S2I scripts

#### assemble
Create an *assemble* script that will build our application, e.g.:
- build python modules
- bundle install ruby gems
- setup application specific configuration

The script can also specify a way to restore any saved artifacts from the previous image.   

#### run
Create a *run* script that will start the application.

#### save-artifacts (optional)
Create a *save-artifacts* script which allows a new build to reuse content from a previous version of the application image.

#### usage (optional)
Create a *usage* script that will print out instructions on how to use the image.

## References
 - https://blog.openshift.com/create-s2i-builder-image/
 - https://docs.openshift.org/latest/creating_images/s2i.html
