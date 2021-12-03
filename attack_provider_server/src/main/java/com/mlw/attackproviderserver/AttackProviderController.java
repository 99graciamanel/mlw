package com.mlw.attackproviderserver;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import java.util.Date;

@RestController
public class AttackProviderController {

    private static final long DELAY = 10000;

    @RequestMapping(path = "/")
    public ResponseEntity<AttackDTO> getAttack() {
        long date = System.currentTimeMillis() + DELAY;
        return ResponseEntity.ok(AttackDTO.builder()
                .date((new Date(date)).toString())
                .dateNs(date * 1000000)
                .ip("10.0.2.5")
                .port("80")
                .build());
    }

}
