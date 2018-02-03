package com.laszloszoboszlai.fabricservice.controller;


import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;


@RestController
@RequestMapping("/fabricservice")
public class FabricController {

    // Initiates the chainCode
    @PostMapping(value = "/init")
    public ResponseEntity<String> initChainCode() {
        return new ResponseEntity<>("", HttpStatus.OK);
    }

    // Returns a random user
    @RequestMapping(name = "/randomuser", method = RequestMethod.GET)
    public ResponseEntity<String> getRandomUser() {
        return new ResponseEntity<>("", HttpStatus.OK);
    }

    // Gets a user data (list of thanks) by key
    @RequestMapping(name = "/getuser", method = RequestMethod.GET)
    public ResponseEntity<String> getUser() {
        return new ResponseEntity<>("", HttpStatus.OK);
    }

    // gets a list of all the available keys in DB
    @RequestMapping(name = "/getkeys", method = RequestMethod.GET)
    public ResponseEntity<String> getKeys() {
        return new ResponseEntity<>("", HttpStatus.OK);
    }

    // adds a thank to the user
    @PostMapping(value = "{userId}/addthank/{value}")
    public ResponseEntity<String> addThank(@PathVariable long userId, @PathVariable long value ) {
        return new ResponseEntity<>("", HttpStatus.OK);
    }

    // adds a user (chaincode needs to be implemented for this)
    @PostMapping(value = "/register/{userID}")
    public ResponseEntity<String> addThank(@RequestBody String user ) {
        return new ResponseEntity<>("", HttpStatus.OK);
    }

}