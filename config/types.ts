// types.ts
export interface Payment {
  id: string;
  amount: number;
  currency: string;
  paymentMethod: PaymentMethod;
  status: PaymentStatus;
  createdAt: Date;
  updatedAt: Date;
}

export enum PaymentMethod {
  CREDIT_CARD = 'credit_card',
  PAYPAL = 'paypal',
  BANK_TRANSFER = 'bank_transfer',
}

export enum PaymentStatus {
  PENDING = 'pending',
  PROCESSING = 'processing',
  SUCCESS = 'success',
  FAILED = 'failed',
}

export interface PaymentGateway {
  name: string;
  config: PaymentGatewayConfig;
}

export interface PaymentGatewayConfig {
  apiKey: string;
  apiSecret: string;
  endpoint: string;
}

export interface PaymentProcessorConfig {
  paymentGateways: PaymentGateway[];
  defaultGateway: string;
}