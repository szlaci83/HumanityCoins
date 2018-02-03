package com.laszloszoboszlai.fabricservice.service;

import com.laszloszoboszlai.fabricservice.repository.FabricThankRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class FabricService {
    @Autowired
    FabricThankRepository thankRepository;

}
