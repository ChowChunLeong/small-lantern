# SOLID Principles Explained Briefly with Problems if Not Followed

## 1️⃣ S - Single Responsibility Principle (SRP)

**Each class/module should have only one reason to change.**

🔴 **Problem if ignored:**  
If a class handles multiple responsibilities, a change in one feature might unintentionally break another, making maintenance difficult.

---

## 2️⃣ O - Open/Closed Principle (OCP)

**Software should be open for extension but closed for modification.**

🔴 **Problem if ignored:**  
Every time a new feature is added, modifying existing code increases the risk of introducing new bugs, making scaling harder.

---

## 3️⃣ L - Liskov Substitution Principle (LSP)

**Subtypes must be substitutable for their base types without altering correctness.**

**Definition**
Liskov Substitution Principle (LSP) is more related to structs (or concrete types)
because it ensures that a subtype (struct) can replace its parent type without breaking functionality.

🔴 **Problem if ignored:**  
If a subclass changes expected behavior, replacing the base type with it might cause unexpected errors, breaking the system.

**Example Violation:**

A Bird class has a Fly() method, but a Penguin subclass cannot fly.
If we substitute Penguin where Bird is expected, it will break the program.
Solution: Instead of forcing Penguin to implement Fly(), we create separate interfaces (FlyingBird and NonFlyingBird).

---

## 4️⃣ I - Interface Segregation Principle (ISP)

**No class should be forced to implement methods it does not use.**

**Definition**
Interface Segregation Principle (ISP) is more related to interfaces,
as it emphasizes having small, focused interfaces instead of forcing a struct to implement unnecessary methods.

🔴 **Problem if ignored:**  
Classes might have unnecessary dependencies on unused methods, leading to bloated and harder-to-maintain code.

**Example Violation:**

An interface Worker has both Work() and Eat().
A Robot class must implement Eat(), even though robots don’t eat.
Solution: Separate interfaces, e.g., Workable for Work() and Eatable for Eat().

---

## 5️⃣ D - Dependency Inversion Principle (DIP)

**Depend on abstractions, not concrete implementations.**

🔴 **Problem if ignored:**  
High-level modules become tightly coupled with low-level implementations, making changes difficult and reducing flexibility.

---

# Efficient Memory Management in Golang

## 📌 Stack vs Heap

| Feature                | Stack                     | Heap                            |
| ---------------------- | ------------------------- | ------------------------------- |
| **Speed**              | Faster                    | Slower                          |
| **Allocation**         | Automatic                 | Manual (via pointer allocation) |
| **Lifetime**           | Limited to function scope | Persists beyond function call   |
| **Size**               | Small                     | Large                           |
| **Garbage Collection** | No                        | Yes                             |

# **Project Structure & Guidelines**

## **Folder Structure**

```
/controller    → Handles API requests
/form         → Structs for request body binding
/database     → Database connection & queries
/model       → Database models
```

## **📌 Form & Controller Naming Convention**

The `** folder file names must align with the **`** folder file names** because only **controller requests** use form structs for JSON binding.

### **Example:**

```
/controller/auth.go   → Handles authentication requests
/form/auth.go        → Contains struct for binding authentication requests
```

This ensures maintainability and consistency.
