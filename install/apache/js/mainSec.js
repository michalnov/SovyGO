let state;
let pki;
let parser = new DOMParser();
const charset = "MNBVCXY123ASDFGHJKL654QWERTZUIOP789qwertzuioplkjhgfdsamnbvcxy";

function setup() {
    pki = forge.pki;
    state = new State();
    let data = state.loadState();
    if (data != false && data != null) {
        aplyState(data);
    } else {
        newState();
    }
}
function sendAESKey(data) {
    let xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            
        }
    };
    xhttp.open("POST", "/key/register/?id=", true);
    xhttp.send(data);
}

function pingSovy() {
    let xhttp = new XMLHttpRequest();
    let data = 
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
        
        }
    };
    xhttp.open("POST", "/key/ping/?id="+state.user.id, true);
    xhttp.send();
}

function importPublicKey(data) {
    let xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var env = new Envelop;
            env.fromEnvelop(this.responseText);
            state.scrypt.serverPub = pki.publicKeyFromPem(env.key);
            reportState("imported");
        } else {
            //alert("failed to import key");
            reportState("failed to import RSA key");
        }
    };
    xhttp.open("GET", "/key/new/?id="+state.user.id, true);
    xhttp.send(); 
}
function encryptRSA(data) {
    return state.scrypt.serverPub.encript(data);
}

function encryptAES(params) {
    
}

function decryptAES(params) {
    
}

function prepareAES(params) {
    
}

function reportState(params) {
    document.getElementById("state").innerHTML = params;
}

function PublicToXml(data) {
    let x = pki.publicKeyToPem(key.publicKey);
    return "<key>"+x+"</key>";
}

function PublicFromXml(data) {
    let xmlDoc = parser.parseFromString(data,"text/xml");
    let y = xmlDoc.getElementsByTagName("key").valu;
    let x = pki.publicKeyFromPem(y);
    return x;
}

function getRandomString(length) {
    let out = "";
    let x = 0;
    for (let i = 0; i < length; i++) {
        out = out + charset[Math.random(charset.length)];
    }
    return out;
}

class State {
    constructor(){
        this.user = new User;
        this.scrypt = new SCrypt;
    }

    loadState(){
        if (typeof(Storage) !== "undefined") {
            let data = localStorage.getItem("sovygo_state");
            if (data == null) {
                
                return false;
            }
            if (data.outdated) {
                return false;
            } else {
                return data;
            }
        } else {
            return false;
        }
    }

    saveState(){
        if (typeof(Storage) !== "undefined") {
            let data = {
                outdated : false,
                user : {
                    username : this.user.name,
                    id : this.user.id
                },
                keys : {
                    //clientPub : pki.publicKeyToPem(this.scrypt.keyClient.publicKey),
                    //clientPri : pki.privateKeyToPem(this.scrypt.keyClient.privateKey),
                    serverPub : pki.publicKeyToPem(this.scrypt.serverPub),
                    symmetricKey : this.scrypt.symmetricKey,
                    symmetricIV : this.scrypt.symmetricIV
                }
            };
            localStorage.setItem();
        }
    }

    clearState(){
        if (typeof(Storage) !== "undefined") {
            localStorage.removeItem("sovygo_state");
        }
    }
}

class User {
    constructor(){
        this.username = "";
        this.id = getRandomString(16);
    }
}

class SCrypt {
    constructor(){
        //this.keyClient = forge.rsa.generateKeyPair();
        this.serverPub = null;
        this.symmetricKey = null;
        this.symmetricIV = null;
    }

    setServerKey(data) {
        this.serverPub = pki.publicKeyFromPem(data);
    }

}

function aplyState(data) {
    state.user.name = data.user.username;
    state.user.id = data.user.id;
    state.scrypt.setServerKey(data.keys.serverPub);
    state.scrypt.symmetricKey = data.keys.symmetricKey;
    state.scrypt.symmetricIV = data.keys.symmetricIV;

}

function newState(params) {
    state.user.id = getRandomString(16);
    state.scrypt.symmetricKey = forge.random.getBytesSync(32);
    importPublicKey();
}
setup();

setInterval(function(){
    pingSovy();
} ,250000);

class Envelop{
    constructor(){
        this.encryption = true;
        this.body = {};
        this.key = "";
        this.check = "";
    }

    toEnvelop(params) {
        this.encryption = false;
    
    }

    encryptToEnvelop(params) {
        this.encryption = true;
        this.key = forge.random.getBytesSync(16);
        let cipher = forge.cipher.createCipher('AES-CBC', state.scrypt.symmetricKey);
        cipher.start({iv: this.key});
        cipher.update(forge.util.createBuffer(params));
        cipher.finish();
        this.body = cipher.output;
    }

    RSAToEnvelop(params) {
        this.encryption = true;
        this.body = state.scrypt.serverPub.encrypt(params);
    }

    fromEnvelop(params) {
        xmlDoc = parser.parseFromString(params, "text/xml");
        this.encryption = xmlDoc.getElementsByTagName("encryption");
        this.body = xmlDoc.getElementsByTagName("Body");
        this.key = xmlDoc.getElementsByTagName("Key");
        this.check = xmlDoc.getElementsByTagName("Check");
        if (this.encryption) {
            let decipher = forge.cipher.createDecipher('AES-CBC', state.scrypt.symmetricKey);
            decipher.start({iv: this.key});
            decipher.update(this.body);
            this.body = decipher.finish();
        }
        return this.body;
    }

    buildEnvelop(){
        let out = "<Data>"+
            "<Head>"+
                "<encryption>"+this.encryption+"</Encryption>"
            +"</Head>"
            "<Body>"+this.body+"</Body>"+
            "<Key>"+this.key+"</Key>"+
            "<Check>"+this.check+"</Check>"
        +"</Data>";
        return out;
    }
}