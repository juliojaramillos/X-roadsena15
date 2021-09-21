package co.gov.ciudadania.models;

import javax.persistence.*;
import java.io.Serializable;

@Entity
@Table( name = "citizen" )
public class Citizen implements Serializable{

    private static final long serialVersionUID = 1L;

    @Id
    private Integer id;

    @Embedded
    private Document document;

    private String name;

    private String lastname;

    private String address;

    @Enumerated( EnumType.STRING )
    private Gender gender;

    private Integer age;


    public Integer getId( ){
        return id;
    }

    public void setId( Integer id ){
        this.id = id;
    }

    public Document getDocument( ){
        return document;
    }

    public void setDocument( Document document ){
        this.document = document;
    }

    public String getName( ){
        return name;
    }

    public void setName( String name ){
        this.name = name;
    }

    public String getLastname( ){
        return lastname;
    }

    public void setLastname( String lastname ){
        this.lastname = lastname;
    }

    public String getAddress( ){
        return address;
    }

    public void setAddress( String address ){
        this.address = address;
    }

    public Gender getGender( ){
        return gender;
    }

    public void setGender( Gender gender ){
        this.gender = gender;
    }

    public Integer getAge( ){
        return age;
    }

    public void setAge( Integer age ){
        this.age = age;
    }
}
