/**
 * @param {number} A
 * @param {number} B
 * @param {number} C
 * @param {number} D
 * @param {number} E
 * @param {number} F
 * @param {number} G
 * @param {number} H
 * @return {number}
 */
var computeArea = function(A, B, C, D, E, F, G, H) {
    var val = (C-A)*(D-B) + (G-E)*(H-F);
    if (E > C || G < A || F > D || H < B) {
        return val;
    }
    val -= (Math.min(C,G) - Math.max(A,E))*(Math.min(D,H) - Math.max(B,F));
    return val;
};