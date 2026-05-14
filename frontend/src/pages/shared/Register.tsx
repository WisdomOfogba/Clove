import React, { useState } from 'react';
import { motion } from 'motion/react';
import { Icon } from '@iconify/react';
import { Link, useNavigate } from 'react-router-dom';
import { Input } from '../../components/ui/Input';
import { Button } from '../../components/ui/Button';
import { api } from '../../lib/api';

export function RegisterPage() {
  const [role, setRole] = useState<'customer' | 'vendor'>('customer');
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const handleRegister = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    const res = await api.register({ role });
    setLoading(false);
    if (res.success) {
       if (role === 'vendor') navigate('/vendor/onboarding'); // Redirect to onboarding
       else navigate('/');
    }
  };

  return (
    <div className="min-h-[85vh] bg-slate-50 flex items-center justify-center p-4 py-12">
      <div className="w-full max-w-md bg-white p-8 sm:p-12 rounded-[3rem] shadow-2xl border border-muted relative overflow-hidden">
        
        <div className="text-center mb-8">
           <h1 className="text-3xl font-black text-dark tracking-tighter">Create Account</h1>
           <p className="text-neutral font-medium mt-1">Join the safest food ecosystem</p>
        </div>

        {/* Role Selector */}
        <div className="flex bg-slate-100 p-1.5 rounded-full mb-8 relative">
           {['customer', 'vendor'].map((r) => (
             <button
               key={r}
               onClick={() => setRole(r as any)}
               className={`flex-1 py-2 text-xs font-black uppercase tracking-widest rounded-full transition-all z-10 ${role === r ? 'text-dark' : 'text-slate-400 hover:text-dark'}`}
             >
               {r}
             </button>
           ))}
           {/* Animated Background */}
           <motion.div 
             className="absolute top-1.5 bottom-1.5 bg-white rounded-full shadow-sm border border-slate-200 pointer-events-none"
             initial={false}
             animate={{
               left: role === 'customer' ? '0.375rem' : '50%',
               width: 'calc(50% - 0.75rem)'
             }}
           />
        </div>

        <form onSubmit={handleRegister} className="space-y-4">
          <Input 
            type="text" 
            placeholder="Full Name" 
            required
            icon={<Icon icon="solar:user-bold-duotone" className="w-5 h-5 text-dark" />}
          />
          <Input 
            type="email" 
            placeholder="Email Address" 
            required
            icon={<Icon icon="solar:letter-bold-duotone" className="w-5 h-5 text-dark" />}
          />
          <Input 
            type="password" 
            placeholder="Password"
            required
            icon={<Icon icon="solar:lock-password-bold-duotone" className="w-5 h-5 text-dark" />}
          />
          
          <Button type="submit" className="w-full h-14 mt-4" disabled={loading}>
            {loading ? <div className="w-6 h-6 border-2 border-white border-t-transparent rounded-full animate-spin" /> : (role === 'vendor' ? 'Start Vendor Onboarding' : 'Create Account')}
          </Button>
        </form>

        <p className="text-center text-sm font-medium text-neutral mt-8 -mb-2">
           Already have an account? <Link to="/login" className="text-dark font-black hover:underline">Sign in</Link>
        </p>

      </div>
    </div>
  );
}
