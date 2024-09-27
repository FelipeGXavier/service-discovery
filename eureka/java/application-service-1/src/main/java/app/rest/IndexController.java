package app.rest;

import org.springframework.cloud.client.discovery.DiscoveryClient;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;

@RestController
public class IndexController {

    private final DiscoveryClient discoveryClient;
    private final RestTemplate restTemplate;

    public IndexController(DiscoveryClient discoveryClient, RestTemplate restTemplate) {
        this.discoveryClient = discoveryClient;
        this.restTemplate = restTemplate;
    }


    @GetMapping(value = "/api/v1/info")
    public String info() {
        return "Ok Service 1";
    }

    @GetMapping(value = "/api/v1/call")
    public String call() {
        var serviceInstance = this.discoveryClient.getInstances("application-service-2")
                .stream()
                .findFirst();
        if (serviceInstance.isEmpty()) {
            return "Error to call service";
        }
        final String requestPath = "/api/v1/info";
        return this.restTemplate.getForEntity(serviceInstance.get().getUri() + requestPath, String.class).getBody();
    }
}
