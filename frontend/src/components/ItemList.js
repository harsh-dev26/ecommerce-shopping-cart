import React, { useState, useEffect } from 'react';
import { getItems, addToCart, getCarts, createOrder, getOrders } from '../services/api';
import './ItemsList.css';

function ItemsList() {
  const [items, setItems] = useState([]);
  const [cartId, setCartId] = useState(null);
  const [showToast, setShowToast] = useState(false);
  const [toastMessage, setToastMessage] = useState('');

  useEffect(() => {
    fetchItems();
  }, []);

  const fetchItems = async () => {
    try {
      const response = await getItems();
      setItems(response.data.data);
    } catch (error) {
      console.error('Error fetching items:', error);
    }
  };

  const handleAddToCart = async (itemId) => {
    try {
      const response = await addToCart([itemId]);
      setCartId(response.data.data.id);
      showToastMessage('Item added to cart!');
    } catch (error) {
      console.error('Error adding to cart:', error);
      window.alert('Failed to add item to cart');
    }
  };

  const handleViewCart = async () => {
    try {
      const response = await getCarts();
      const userCarts = response.data.data;
      
      if (userCarts.length === 0) {
        window.alert('Your cart is empty');
        return;
      }

      const cartDetails = userCarts.map(cart => {
        const items = cart.items ? cart.items.map(item => 
          `Item ID: ${item.id}, Name: ${item.name}`
        ).join('\n') : 'No items';
        
        return `Cart ID: ${cart.id}\n${items}`;
      }).join('\n\n');

      window.alert(cartDetails);
    } catch (error) {
      console.error('Error fetching cart:', error);
      window.alert('Failed to fetch cart');
    }
  };

  const handleViewOrders = async () => {
    try {
      const response = await getOrders();
      const orders = response.data.data;
      
      if (orders.length === 0) {
        window.alert('No orders yet');
        return;
      }

      const orderDetails = orders.map(order => 
        `Order ID: ${order.id}, Cart ID: ${order.cart_id}`
      ).join('\n');

      window.alert(orderDetails);
    } catch (error) {
      console.error('Error fetching orders:', error);
      window.alert('Failed to fetch orders');
    }
  };

  const handleCheckout = async () => {
    if (!cartId) {
      window.alert('Your cart is empty. Add items before checkout.');
      return;
    }

    try {
      await createOrder(cartId);
      setCartId(null);
      showToastMessage('Order successful!');
    } catch (error) {
      console.error('Error creating order:', error);
      window.alert('Failed to create order');
    }
  };

  const showToastMessage = (message) => {
    setToastMessage(message);
    setShowToast(true);
    setTimeout(() => setShowToast(false), 3000);
  };

  const handleLogout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('user_id');
    window.location.reload();
  };

  return (
    <div className="items-container">
      <div className="header">
        <h1>E-Commerce Store</h1>
        <div className="header-buttons">
          <button onClick={handleViewCart} className="btn btn-secondary">
            View Cart
          </button>
          <button onClick={handleViewOrders} className="btn btn-secondary">
            Order History
          </button>
          <button onClick={handleCheckout} className="btn btn-primary">
            Checkout
          </button>
          <button onClick={handleLogout} className="btn btn-danger">
            Logout
          </button>
        </div>
      </div>

      <div className="items-grid">
        {items.map((item) => (
          <div key={item.id} className="item-card">
            <h3>{item.name}</h3>
            <p>Status: {item.status}</p>
            <button 
              onClick={() => handleAddToCart(item.id)}
              className="btn btn-add"
            >
              Add to Cart
            </button>
          </div>
        ))}
      </div>

      {showToast && (
        <div className="toast">
          {toastMessage}
        </div>
      )}
    </div>
  );
}

export default ItemsList;