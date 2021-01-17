/*
* https://www.codewars.com/kata/547f1a8d4a437abdf800055c
* So the purpose of this kata is to create an object with accessible and "updatable" (can i say that?) properties.
*
* If .firstName or .lastName are changed, then .fullName should also be changed
* If .fullName is changed, then .firstName and .lastName should also be changed.
* Note : "input format" to .fullName will be firstName + space+ lastName. If given fullName isn't valid then no property is changed.
* */
function NamedOne(first, last) {
    this.firstName = first;
    this.lastName = last;
    Object.defineProperty(this, "fullName", {
        get: function () {
            return this.firstName + ' ' + this.lastName;
        },
        set: function (text) {
            if(text.split(' ').length < 2){
                return;
            }
            this.firstName = text.split(' ')[0];
            this.lastName = text.split(' ')[1];
        }
    });
}
