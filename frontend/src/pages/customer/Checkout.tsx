import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { motion, AnimatePresence } from 'motion/react';
import { Icon } from '@iconify/react';
import { useCart } from '../../contexts/CartContext';
import { Button } from '../../components/ui/Button';
import { Input } from '../../components/ui/Input';
import { api } from '../../lib/api';

export function CheckoutPage() {
  const { totalPrice, clearCart, items } = useCart();
  const navigate = useNavigate();
  const [showSquadModal, setShowSquadModal] = useState(false);
  const [isProcessing, setIsProcessing] = useState(false);

  const totalAmount = totalPrice + 1500; // Adding delivery fee

  const handlePaymentSuccess = async () => {
    setIsProcessing(true);
    // Simulate creating the order
    const res = await api.createOrder({ items, total: totalAmount });
    if (res.success) {
      clearCart();
      navigate('/order-confirmation');
    }
  };

  return (
    <div className="container mx-auto px-4 py-12 max-w-5xl">
      <h1 className="text-4xl font-black text-dark mb-10">Checkout</h1>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <div className="lg:col-span-2 space-y-8">
          
          <div className="bg-white p-8 rounded-[2.5rem] border border-muted shadow-sm">
            <h2 className="text-2xl font-black text-dark mb-6 flex items-center gap-2">
              <Icon icon="solar:map-point-bold-duotone" className="text-primary"/> Delivery Details
            </h2>
            <div className="space-y-4">
              <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                <Input placeholder="First Name" />
                <Input placeholder="Last Name" />
              </div>
              <Input placeholder="Phone Number" type="tel" />
              <Input placeholder="Full Delivery Address" />
              <textarea 
                placeholder="Delivery Notes (Optional)" 
                className="w-full rounded-2xl border-2 border-muted bg-white p-5 text-sm font-bold text-dark focus:border-primary outline-none resize-none h-32"
              ></textarea>
            </div>
          </div>

          <div className="bg-white p-8 rounded-[2.5rem] border border-muted shadow-sm flex items-center justify-between">
            <div>
              <h2 className="text-2xl font-black text-dark flex items-center gap-2">
                <Icon icon="solar:shield-check-bold" className="text-primary"/> Security Guarantee
              </h2>
              <p className="text-neutral font-medium mt-2">Your delivery bag will be sealed and strictly geo-locked.</p>
            </div>
            <Icon icon="solar:lock-password-bold-duotone" className="w-16 h-16 text-slate-100" />
          </div>

        </div>

        <div>
          <div className="bg-slate-50 border border-muted p-8 rounded-[2.5rem] sticky top-24">
            <h3 className="text-xl font-black text-dark mb-6">Payment</h3>
            
            <div className="py-4 border-y border-muted mb-6 space-y-2 text-sm font-bold text-neutral">
              <div className="flex justify-between"><span>Subtotal</span><span>₦{totalPrice}</span></div>
              <div className="flex justify-between"><span>Delivery</span><span>₦1,500</span></div>
            </div>

            <div className="flex justify-between items-center mb-8">
              <span className="font-bold text-dark text-lg">Total to Pay</span>
              <span className="text-3xl font-black text-primary">₦{totalAmount}</span>
            </div>

            <Button 
              className="w-full h-16 text-lg tracking-wide rounded-2xl shadow-[0_10px_20px_rgba(0,0,0,0.1)]"
              onClick={() => setShowSquadModal(true)}
            >
              Pay with Squadco
            </Button>
            <p className="text-[10px] uppercase tracking-widest text-neutral text-center mt-4 font-bold flex items-center justify-center gap-1">
              <Icon icon="solar:lock-keyhole-bold" /> Secured by Squad
            </p>
          </div>
        </div>
      </div>

      {/* Simulated Squadco Modal */}
      <AnimatePresence>
        {showSquadModal && (
          <div className="fixed inset-0 z-50 flex items-center justify-center p-4">
            <motion.div 
              initial={{ opacity: 0 }} animate={{ opacity: 1 }} exit={{ opacity: 0 }}
              className="absolute inset-0 bg-dark/60 backdrop-blur-sm"
              onClick={() => !isProcessing && setShowSquadModal(false)}
            />
            <motion.div 
              initial={{ scale: 0.9, opacity: 0 }} animate={{ scale: 1, opacity: 1 }} exit={{ scale: 0.9, opacity: 0 }}
              className="bg-white w-full max-w-sm rounded-[2rem] p-8 relative z-10 text-center shadow-2xl"
            >
              <div className="w-16 h-16 bg-[#ffeed9] text-[#e85d04] rounded-full flex items-center justify-center mx-auto mb-6">
                 <Icon icon="solar:wallet-money-bold-duotone" className="w-8 h-8" />
              </div>
              <h3 className="text-2xl font-black text-dark mb-2">Squad Checkout</h3>
              <p className="text-neutral font-medium mb-8">Testing environment. Click below to simulate payment.</p>

              {isProcessing ? (
                <div className="flex flex-col items-center gap-4">
                  <div className="w-8 h-8 border-4 border-[#e85d04] border-t-transparent rounded-full animate-spin"></div>
                  <p className="font-bold text-neutral">Processing payment...</p>
                </div>
              ) : (
                <Button 
                  className="w-full bg-[#e85d04] text-white hover:bg-[#d05303] h-14"
                  onClick={handlePaymentSuccess}
                >
                  Simulate Success ₦{totalAmount}
                </Button>
              )}
            </motion.div>
          </div>
        )}
      </AnimatePresence>

    </div>
  );
}
