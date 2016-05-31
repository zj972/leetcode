/**
 * @param {string} secret
 * @param {string} guess
 * @return {string}
 */
//Distribution 67.53%,runtime 152ms
var getHint = function(secret, guess) {
    var bulls = 0;
    var cows = 0;
    var numbers = [0,0,0,0,0,0,0,0,0,0];
    for (var i = 0; i < secret.length; i++) {
        if (secret[i] == guess[i]) bulls++;
        else {
            if (numbers[secret[i]-"0"]++ < 0) cows++;
            if (numbers[guess[i]-"0"]-- > 0) cows++;
        }
    }
    return bulls + "A" + cows + "B";
};
//旧版本，效率低
//Distribution 2.60%,runtime 304ms
var getHint = function(secret, guess) {
    var bull = 0;
    var cow = 0;
    secret = secret.split("");
    guess = guess.split("");
    for(var i = 0, n = secret.length; i < n; i++){
        if(secret[i] === guess[i]){
            bull++;
            secret[i] = null;
            guess[i] = null;
        }
    }
    for(var j = 0, m = guess.length; j < m; j++){
        if(guess[j] === null) continue;
        else{
            for(var k = 0; k < n; k++){
                if(guess[j] === secret[k]){
                    cow++;
                    guess[j] = null;
                    secret[k] = null;
                    break;
                }
            }
        }
    }
    return bull+"A"+cow+"B";
};