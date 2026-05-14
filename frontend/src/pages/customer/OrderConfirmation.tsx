import React, { useEffect } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import { motion } from 'motion/react';
import { Icon } from '@iconify/react';
import { Button } from '../../components/ui/Button';

export function OrderConfirmationPage() {
  const navigate = useNavigate();

  // Scroll to top
  useEffect(() => {
    window.scrollTo(0, 0);
  }, []);

  return (
    <div className="min-h-[80vh] flex items-center justify-center container mx-auto px-4 py-20">
      <motion.div 
        initial={{ opacity: 0, scale: 0.9 }}
        animate={{ opacity: 1, scale: 1 }}
        className="bg-white max-w-lg w-full p-10 rounded-[3rem] shadow-[0_20px_60px_-15px_rgba(0,0,0,0.1)] text-center border border-muted"
      >
        <motion.div 
          initial={{ scale: 0 }}
          animate={{ scale: 1 }}
          transition={{ type: "spring", delay: 0.2 }}
          className="w-24 h-24 bg-green-100 text-green-500 rounded-full flex items-center justify-center mx-auto mb-8 shadow-inner"
        >
          <Icon icon="solar:check-circle-bold" className="w-12 h-12" />
        </motion.div>
        
        <h1 className="text-4xl font-black text-dark mb-4">Payment Successful!</h1>
        <p className="text-neutral font-medium text-lg mb-2">Your order has been placed and securely sent to the vendor.</p>
        <p className="text-sm font-bold text-slate-400 mb-10">Order ID: #{Math.floor(Math.random() * 100000)}</p>

        <div className="bg-slate-50 p-6 rounded-3xl text-left border border-muted mb-8 relative overflow-hidden">
          <div className="absolute top-0 right-0 w-32 h-32 bg-primary/10 rounded-full blur-3xl" />
          <h3 className="font-black text-dark mb-2 flex items-center gap-2">
            <Icon icon="solar:shield-warning-bold-duotone" className="text-primary" />
            Security Locked
          </h3>
          <p className="text-sm font-medium text-neutral">
            The kitchen is currently being live-audited for your meal preparation. A courier will be assigned shortly.
          </p>
        </div>

        <div className="flex flex-col gap-4">
          <Button onClick={() => navigate('/orders')} className="h-14">
            Track My Order
          </Button>
          <Link to="/">
            <Button variant="ghost" className="w-full text-neutral">
              Return to Home
            </Button>
          </Link>
        </div>
      </motion.div>
    </div>
  );
}
