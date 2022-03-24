package com.example.facadepatternexample;

import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
@Transactional
public class UserMobileCertificationService {
    private final UserMobileCertificationFacade userMobileCertificationFacade;

    public UserMobileCertificationService(UserStore userStore, SecretCodeStore secretCodeStore) {
        this.userMobileCertificationFacade = new UserMobileCertificationFacade(userStore, secretCodeStore);
    }

    public void startCertification(String mobile) {
        userMobileCertificationFacade.startCertification(mobile);
    }

    public void confirmCertification(String mobile, String code) {
        userMobileCertificationFacade.confirmCertification(mobile, code);
    }
}
