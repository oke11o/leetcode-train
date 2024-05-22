/*
  ДЛЯ ИНФОРМАЦИИ

  - другие не видят, что ты делаешь в этом Pen'e;
  - новый код не меняет шаблон для остальных;
  - при перезагрузке страницы все сброситься;
  - форкай, если хочешь сохранить

  УДАЧНОГО КОДИНГА
*/

class ATM {
    vault = {
        5000: 0,
        2000: 0,
        1000: 0,
        500: 0,
        100: 0,
        50: 0,
    };

    deposit(bills) {
        if (!bills || !bills.length) {
            return Error(" Положите деньги в купюроприемник");
        }
        let acceptedSum = 0;
        let returnedBills = [];
        for (let bill of bills) {
            if (bill in this.vault) {
                this.vault[bill]++;
                acceptedSum += bill;
            } else {
                returnedBills.push(bill);
            }
        }
        return {acceptedSum:acceptedSum, returnedBills:returnedBills};
    }

    whithdrow(amount) {
        if (amount === 0) {
            return Error("Укажите корректную сумму");
        }
        if (amount > this.total) {
            return Error("Не могу выдать нужную сумму, недостаточно средств");
        }
        let acceptReverse = this.accept.reverse();
        const result = []; // [1000]
        // while (amount > 0) { // 1100
        //     let abort = true;
        //     for (let bill of acceptReverse) { // {5000:2,2000:0,1000:1,500:2,100:1,50:2}
        //         if (bill <= amount && this.vault[bill] > 0) { //1000 < 2100
        //             result.push(bill);
        //             this.vault[bill]--;
        //             amount -= bill;
        //             abort = false;
        //             break;
        //         }
        //     }
        //     if (abort) {
        //         return Error('Error: Не могу выдать нужную сумму, недостаточно купюр')
        //     }
        // }

        for (let bill of acceptReverse) { // {5000:2,2000:0,1000:1,500:1,100:1,50:2}
            while (bill <= amount && amount > 0 && this.vault[bill] > 0) {
                result.push(bill);
                this.vault[bill]--;
                amount -= bill;
            }
        }
        if (amount > 0) {
            return Error('Error: Не могу выдать нужную сумму, недостаточно купюр')
        }
        return result;
    }

    // возвращает массив купюр который доступен на прием/выдачу
    get accept() {
        return Object.keys(this.vault);
    }

    // возвращает сколько всего денег во внутреннем хранилище
    get total() {
        let sum = 0;
        for (let bill in this.vault) {
            sum += bill * this.vault[bill];
        }
        return sum;
    }

    // возвращает касету с деньгами в виде объекта при инкассации
    get collect() {
        return this.vault;
    }
}

const atm = new ATM();
console.log(atm.accept); // [ 50, 100, 500, 1000, 2000, 5000 ]
console.log(atm.whithdrow(3500)); // Error: Не могу выдать нужную сумму, недостаточно средств
console.log(atm.deposit([])); // Error: Положите деньги в купюроприемник
console.log(atm.deposit([5000, 1000, 5000, 500, 100, 50, 50])); // Внесено 11700
let data = atm.deposit([500, 10, 5]);
console.log(`Внесено ${data.acceptedSum}, Заберите нераспознанные купюры [${data.returnedBills}]`); // Внесено 500, Заберите нераспознанные купюры [10, 5]
// console.log(atm.whithdrow(3500)); // Error: Не могу выдать нужную сумму, недостаточно купюр
console.log(atm.whithdrow(2100)); // [1000, 500, 500, 100]
console.log(atm.whithdrow(0)); // Error: Укажите корректную сумму
console.log(atm.total); //10100
console.log(atm.collect); // { '50': 2, '100': 0, '500': 0, '1000': 0, '2000': 0, '5000': 2 }
