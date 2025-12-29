from locust import HttpUser, task

class HelloWorldUser(HttpUser):
    @task
    def hello_world(self):
        self.client.get("http://dice:8080/rolldice")
        self.client.get("http://dice:8080/rolldice/Alice")