package com.example.facadepatternexample;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.EnableAspectJAutoProxy;

@SpringBootApplication
public class FacadePatternExampleApplication {

    public static void main(String[] args) {
        SpringApplication.run(FacadePatternExampleApplication.class, args);

    }

}
