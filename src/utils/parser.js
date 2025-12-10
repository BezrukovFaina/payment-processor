const fs = require('fs');
const csv = require('csv-parser');

class PaymentParser {
  constructor(filePath) {
    this.filePath = filePath;
  }

  async parsePayments() {
    const payments = [];
    const fileStream = fs.createReadStream(this.filePath);
    return new Promise((resolve, reject) => {
      fileStream
        .pipe(csv())
        .on('data', (data) => {
          const payment = {
            id: data.id,
            amount: parseFloat(data.amount),
            type: data.type,
            timestamp: new Date(data.timestamp),
          };
          payments.push(payment);
        })
        .on('end', () => {
          resolve(payments);
        })
        .on('error', (error) => {
          reject(error);
        });
    });
  }
}

module.exports = PaymentParser;