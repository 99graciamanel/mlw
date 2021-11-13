package com.mlw.attackproviderserver;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import java.util.Date;

@RestController
public class AttackProviderController {

    private static final long DELAY = 3600000;

    @RequestMapping(path = "/")
    public ResponseEntity<AttackDTO> getAttack() {
        return ResponseEntity.ok(AttackDTO.builder()
                .date((new Date(System.currentTimeMillis() + DELAY)).toString())
                .ip("10.0.2.15")
                .build());
    }

}
