
package com.example.springboot;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.http.ResponseEntity;
import org.springframework.http.HttpStatus;

@RestController
public class HelloController {

	@RequestMapping("/")
	public String index() {
		return "Hello World!";
	}

	@RequestMapping("/health")
	public ResponseEntity<String> health() {
		String message = "OK!";
		return new ResponseEntity<>(message, HttpStatus.OK);
	}

}

