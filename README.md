
# fga-golang-04
NOTED:
Tidak bisa deploy kode ke railway-
Status sudah terhubung ke github, tapi masih limited trial.

Pesan:
To prevent abuse of free resources, we check GitHub accounts before enabling code deployments on Railway.

Unfortunately, the GitHub account unta-terbang didn't give us enough information to let you deploy code (Learn More).Please continue with a Limited Trial or upgrade your account below.

DAFTAR API LOKAL
-
Users

-------------------------------------------------------------------------------------
[PUBLIC]    PUSH    http://localhost:8080/users/register

[PUBLIC]    PUSH    http://localhost:8080/users/login

[Aunthentication & Authorization]    PUT    http://localhost:8080/users/login

[Aunthentication & Authorization]    DELETE    http://localhost:8080/users/login


Photos
-------------------------------------------------------------------------------------
[Aunthentication]    PUSH    http://localhost:8080/photos

[Aunthentication]    GET    http://localhost:8080/photos/

[Aunthentication]    GET    http://localhost:8080/photos/:id

[Aunthentication & Authorization]    PUT    http://localhost:8080/photos/:id

[Aunthentication & Authorization]    DELETE    http://localhost:8080/photos/:id


Comments
-------------------------------------------------------------------------------------
[Aunthentication]   PUSH    http://localhost:8080/comments

[Aunthentication]   GET    http://localhost:8080/comments/

[Aunthentication]   GET   http://localhost:8080/comments/:id

[Aunthentication & Authorization]    PUT    http://localhost:8080/comments/:id

[Aunthentication & Authorization]    DELETE    http://localhost:8080/comments/:id


Social Medias
-------------------------------------------------------------------------------------
[Aunthentication]    PUSH    http://localhost:8080/socialmedias

[Aunthentication & Authorization]    GET    http://localhost:8080/socialmedias/

[Aunthentication & Authorization]    GET    http://localhost:8080/socialmedias/:id

[Aunthentication & Authorization]    PUT    http://localhost:8080/socialmedias/:id

[Aunthentication & Authorization]    DELETE    http://localhost:8080/socialmedias/:id
