package com.example.facadepatternexample;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;
import javax.persistence.Table;

@Entity
@Table(name="SECRET_CODE")
public class SecretCode {

    public SecretCode() {}

    public SecretCode(String code) {
        this.code = code;
    }

    @Id
    private Long id;

    private Boolean isLock = false;

    private String code;

    public void setId(Long id) {
        this.id = id;
    }

    public String getCode() {
        if (this.isLock) throw new RuntimeException("잠겨진 코드임 접근 불가함");
        return this.code;
    }

    public void Lock() {
        this.isLock = true;
    }

    public Boolean compare(String plain) {
        return code.equals(plain);
    }
}
