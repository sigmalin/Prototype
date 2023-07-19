## Prototype - ssl

<br><br>

# 介紹

- 建立 自簽的 SSL 

<br><br>

# 手順

+ 安裝 openssl

<pre>
> sudo apt-get update
> sudo apt-get install openssl
</pre>

安裝完後可以使用下述指令檢查是否安裝成功

<pre>
> openssl version
OpenSSL 3.0.2 15 Mar 2022 (Library: OpenSSL 3.0.2 15 Mar 2022)
</pre>


+ 產生 Private Key

<pre>
> openssl genrsa -aes128 -out server.key 2048
Enter PEM pass phrase:
Verifying - Enter PEM pass phrase:
</pre>

需輸入最少4個字元的密碼

但因為 golang 函式庫的 LoadX509KeyPair 並沒有傳入密碼的選項，需將密碼移除才能使用

<pre>
> openssl rsa -in server.key -out server.key
Enter pass phrase for server.key:
writing RSA key
</pre>


+ 產生 CSR / Certificate Signing Request

<pre>
> openssl req -new -days 3650 -key server.key -out server.csr
Ignoring -days without -x509; not generating a certificate
Enter pass phrase for server.key:
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) [AU]:TW
State or Province Name (full name) [Some-State]:sigma
Locality Name (eg, city) []:
Organization Name (eg, company) [Internet Widgits Pty Ltd]:
Organizational Unit Name (eg, section) []:
Common Name (e.g. server FQDN or YOUR name) []:
Email Address []:

Please enter the following 'extra' attributes
to be sent with your certificate request
A challenge password []:
An optional company name []:
</pre>


+ 產生 Certification file

<pre>
> openssl x509 -in server.csr -out server.pem -req -signkey server.key -days 3650
</pre>