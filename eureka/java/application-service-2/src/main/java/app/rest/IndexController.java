package app.rest;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class IndexController {

    @GetMapping(value = "/api/v1/info")
    public String info() {
        return "Ok Service 2";
    }
}
