package com.example.facadepatternexample;

public class UserMobileCertificationFacade {
    private final UserStore userSystem;

    private final SecretCodeStore secretCodeSystem;

    private final SMSSender smsSystem = new SMSSender();

    public UserMobileCertificationFacade(UserStore userSystem, SecretCodeStore secretCodeSystem) {
        this.userSystem = userSystem;
        this.secretCodeSystem = secretCodeSystem;
    }


    public void startCertification(String mobile) {
        User user;
        try {
            user = userSystem.findByMobile(mobile);
        } catch (Exception e) {
            user = userSystem.newUser(mobile);
        }

        var code = secretCodeSystem.genCode(user);
        smsSystem.sendSMS(mobile, code.getCode());
        code.Lock();
    }

    public void confirmCertification(String mobile, String secretCode) {
        var user = userSystem.findByMobile(mobile);
        if (secretCodeSystem.Find(user).compare(secretCode)) {
            user.Activate();
        }
    }
}
