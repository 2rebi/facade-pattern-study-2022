package com.example.facadepatternexample;


import javax.persistence.*;

@Entity
@Table(name="USER")
public class User {
    @Id
    @GeneratedValue
    private Long id;

    private Boolean active = false;

    public String mobile;

    public Long getId() {
        return this.id;
    }

    public Boolean isActive() {
        return this.active;
    }

    public void Activate() {
        this.active = true;
    }

}
