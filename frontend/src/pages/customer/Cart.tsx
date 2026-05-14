import React from 'react';
import { motion } from 'motion/react';
import { Icon } from '@iconify/react';
import { Link, useNavigate } from 'react-router-dom';
import { useCart } from '../../contexts/CartContext';
import { Button } from '../../components/ui/Button';

export function CartPage() {
  const { items, removeItem, totalPrice, totalItems } = useCart();
  const navigate = useNavigate();

  if (totalItems === 0) {
    return (
      <div className="min-h-[70vh] flex flex-col items-center justify-center p-4">
        <div className="w-32 h-32 bg-slate-100 rounded-full flex items-center justify-center mb-6">
          <Icon icon="solar:cart-large-2-bold-duotone" className="w-16 h-16 text-slate-300" />
        </div>
        <h2 className="text-3xl font-black text-dark mb-2">Your cart is empty</h2>
        <p className="text-neutral font-medium mb-8">Looks like you haven't added anything yet.</p>
        <Link to="/restaurants">
          <Button>Explore Kitchens</Button>
        </Link>
      </div>
    );
  }

  return (
    <div className="container mx-auto px-4 py-12 max-w-4xl">
      <h1 className="text-4xl font-black text-dark mb-8">Your Cart</h1>

      <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
        <div className="md:col-span-2 space-y-4">
          {items.map((item) => (
            <motion.div 
              layout
              key={item.id} 
              className="bg-white p-4 sm:p-6 rounded-[2rem] border border-muted flex items-center justify-between gap-4"
            >
              <div>
                <p className="text-xs font-bold text-neutral uppercase tracking-widest mb-1">{item.vendorName}</p>
                <h3 className="text-lg font-black text-dark">{item.name}</h3>
                <p className="text-primary font-black mt-2">₦{item.price}</p>
              </div>
              
              <div className="flex items-center gap-4">
                <div className="bg-slate-50 px-4 py-2 rounded-xl font-bold border border-muted">
                  x {item.quantity}
                </div>
                <button 
                  onClick={() => removeItem(item.id)}
                  className="w-10 h-10 rounded-full bg-red-50 text-red-500 hover:bg-red-100 flex items-center justify-center transition-colors"
                >
                  <Icon icon="solar:trash-bin-trash-bold" />
                </button>
              </div>
            </motion.div>
          ))}
        </div>

        <div className="bg-dark text-white rounded-[2rem] p-8 h-fit sticky top-24">
          <h3 className="text-xl font-black mb-6 border-b border-white/10 pb-4">Order Summary</h3>
          <div className="space-y-4 mb-6 text-sm font-medium text-white/70">
            <div className="flex justify-between"><span>Subtotal ({totalItems} items)</span><span>₦{totalPrice}</span></div>
            <div className="flex justify-between"><span>Delivery Fee</span><span>₦1,500</span></div>
          </div>
          <div className="pt-4 border-t border-white/10 flex justify-between items-center mb-8">
            <span className="font-bold text-white/50">Total</span>
            <span className="text-3xl font-black text-white">₦{totalPrice + 1500}</span>
          </div>
          
          <Button 
            className="w-full bg-primary text-dark hover:bg-accent" 
            onClick={() => navigate('/checkout')}
          >
            Checkout Securely
          </Button>
        </div>
      </div>
    </div>
  );
}
