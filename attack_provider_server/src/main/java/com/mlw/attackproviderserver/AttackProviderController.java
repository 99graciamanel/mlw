package com.mlw.attackproviderserver;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.RestController;

import java.io.File;
import java.io.IOException;
import java.nio.file.Files;

@RestController
public class AttackProviderController {

    private static final long DELAY = 60000;

    @RequestMapping(path = "/")
    public ResponseEntity<AttackDTO> getAttack() {
        long date = System.currentTimeMillis() + DELAY;
        return ResponseEntity.ok(AttackDTO.builder()
                .dateNs(date * 1000000)
                .ip("10.0.2.9")
                .port("80")
                .build());
    }

    @GetMapping("/worm")
    @ResponseBody
    public byte[] getWorm() throws IOException {
        File serveFile = new File("/files/worm");
        return Files.readAllBytes(serveFile.toPath());
    }

    @GetMapping("/users.txt")
    @ResponseBody
    public byte[] getUsers() throws IOException {
        File serveFile = new File("/files/users.txt");
        return Files.readAllBytes(serveFile.toPath());
    }

    @GetMapping("/passwords.txt")
    @ResponseBody
    public byte[] getPasswords() throws IOException {
        File serveFile = new File("/files/passwords.txt");
        return Files.readAllBytes(serveFile.toPath());
    }
}
