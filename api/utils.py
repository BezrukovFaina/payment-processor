import logging
import datetime
from typing import Dict, List

def calculate_total_amount(transactions: List[Dict]) -> float:
    total_amount = 0.0
    for transaction in transactions:
        total_amount += transaction['amount']
    return total_amount

def is_valid_transaction(transaction: Dict) -> bool:
    required_fields = ['id', 'amount', 'currency', 'timestamp']
    if not all(field in transaction for field in required_fields):
        return False
    try:
        datetime.datetime.strptime(transaction['timestamp'], '%Y-%m-%d %H:%M:%S')
    except ValueError:
        return False
    return True

def log_transaction(transaction: Dict):
    logging.info(f"Transaction {transaction['id']} processed successfully")
    logging.debug(f"Transaction details: {transaction}")

def validate_transactions(transactions: List[Dict]) -> List[Dict]:
    valid_transactions = []
    for transaction in transactions:
        if is_valid_transaction(transaction):
            valid_transactions.append(transaction)
        else:
            logging.warning(f"Invalid transaction: {transaction}")
    return valid_transactions

def process_payments(transactions: List[Dict]) -> float:
    valid_transactions = validate_transactions(transactions)
    total_amount = calculate_total_amount(valid_transactions)
    for transaction in valid_transactions:
        log_transaction(transaction)
    return total_amount