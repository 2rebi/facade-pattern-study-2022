package com.example.facadepatternexample;

import org.springframework.stereotype.Repository;

import javax.persistence.EntityManager;
import javax.persistence.PersistenceContext;

@Repository
public class SecretCodeStore {
    @PersistenceContext
    EntityManager em;

    public SecretCode genCode(User user) {
        var code = new SecretCode("1234");
        code.setId(user.getId());
        em.merge(code);
        return code;
    }

    public SecretCode Find(User user) {
        return em.find(SecretCode.class, user.getId());
    }
}
