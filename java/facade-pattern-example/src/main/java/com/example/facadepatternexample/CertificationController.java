package com.example.facadepatternexample;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class CertificationController {

    private final UserMobileCertificationService service;

    public CertificationController(UserMobileCertificationService service) {
        this.service = service;
    }

    static class InputStartCertification {
        public String mobile;
    }

    @PostMapping("/start-certification")
    public void startCertification(@RequestBody InputStartCertification body) {
        service.startCertification(body.mobile);
    }

    static class InputConfirmCertification {
        public String mobile;
        public String code;
    }

    @PostMapping("/confirm-certification")
    public void confirmCertification(@RequestBody InputConfirmCertification body) {
        service.confirmCertification(body.mobile, body.code);
    }
}
