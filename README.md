# APMC Billing App Backend  

## Overview  
The **APMC Billing App Backend** is designed to manage billing operations for APMC traders. It provides RESTful APIs to handle user accounts, consignees, receivers, invoices, products, and payment statuses. Additionally, the system generates invoices in PDF format for easy download and record-keeping.  

## Key Features  
- **User Management:** Create, login, and delete user accounts.  
- **Consignee & Receiver Management:** Add, retrieve, and delete consignees and receivers.  
- **Invoice Handling:** Create invoices, add products, retrieve invoice details, and download PDFs.  
- **Product Management:** Add, view, and delete products linked to specific invoices.  
- **Payment Management:** Update payment status once transactions are complete.  

## API Endpoints  

### 🔑 User Management  
- **Create User:** `POST /create/user`  
- **Login User:** `POST /login`  
- **Delete User:** `DELETE /delete/<USER_ID>`  

### 🚚 Consignee Management  
- **Create Consignee:** `POST /create/consignee`  
- **Get All Consignees:** `GET /get/consignees/<USER_ID>`  
- **Delete Consignee:** `DELETE /delete/consignee/<CONSIGNEE_ID>`  

### 📦 Receiver Management  
- **Create Receiver:** `POST /create/receiver`  
- **Get All Receivers:** `GET /get/receivers/<USER_ID>`  
- **Delete Receiver:** `DELETE /delete/receiver/<RECEIVER_ID>`  

### 📄 Invoice Management  
- **Create Invoice:** `POST /create/invoice`  
- **Get All Invoices:** `GET /get/invoices/<USER_ID>`  
- **Download Invoice (PDF):** `GET /download/invoice/<INVOICE_ID>`  
- **Delete Invoice:** `DELETE /delete/invoice/<INVOICE_ID>`  
- **Update Payment Status:** `PATCH /update/invoice/payment/<INVOICE_ID>`  

### 🛒 Product Management  
- **Create Product:** `POST /create/product`  
- **Get Products for Invoice:** `GET /get/products/<INVOICE_ID>`  
- **Delete Product:** `DELETE /delete/product/<PRODUCT_ID>`  

## Invoice PDF Generation  
After adding consignees, receivers, products, and invoices, the app generates a **PDF invoice** that can be downloaded for records.  
