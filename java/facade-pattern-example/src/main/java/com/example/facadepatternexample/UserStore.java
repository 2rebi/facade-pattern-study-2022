package com.example.facadepatternexample;

import org.springframework.stereotype.Repository;

import javax.persistence.EntityManager;
import javax.persistence.PersistenceContext;

@Repository
public class UserStore {

    @PersistenceContext
    EntityManager em;

    public User newUser(String mobile) {
        var user = new User();
        user.mobile = mobile;
        em.persist(user);
        return user;
    }

    public User find(Long id) {
        return em.find(User.class, id);
    }

    public User findByMobile(String mobile) {
        return em.createQuery("select u from User u where u.mobile = :mobile", User.class)
                .setParameter("mobile", mobile)
                .getSingleResult();
    }
}
