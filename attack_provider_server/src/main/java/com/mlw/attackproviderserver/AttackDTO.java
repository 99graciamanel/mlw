package com.mlw.attackproviderserver;

import lombok.Builder;
import lombok.Data;

import javax.xml.bind.annotation.XmlElement;
import javax.xml.bind.annotation.XmlRootElement;
import java.io.Serializable;

@Data
@Builder
@XmlRootElement(name = "AttackDTO")
public class AttackDTO implements Serializable {

    @XmlElement(name = "date", required = true)
    private String date;

    @XmlElement(name = "dateNs", required = true)
    private long dateNs;

    @XmlElement(name = "ip", required = true)
    private String ip;

    @XmlElement(name = "port", required = true)
    private String port;
}
