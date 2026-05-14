import React, { useState } from 'react';
import { motion } from 'motion/react';
import { Icon } from '@iconify/react';
import { Link, useNavigate } from 'react-router-dom';
import { Input } from '../../components/ui/Input';
import { Button } from '../../components/ui/Button';
import { api } from '../../lib/api';

export function LoginPage() {
  const [role, setRole] = useState<'customer' | 'vendor' | 'admin'>('customer');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    const res = await api.login({ email, password, role });
    setLoading(false);
    if (res.token) {
       if (role === 'vendor') navigate('/vendor/dashboard');
       else if (role === 'admin') navigate('/admin/dashboard');
       else navigate('/');
    }
  };

  return (
    <div className="min-h-[85vh] bg-slate-50 flex items-center justify-center p-4">
      <div className="w-full max-w-md bg-white p-8 sm:p-12 rounded-[3rem] shadow-2xl border border-muted relative overflow-hidden">
        
        {/* Top styling */}
        <div className="absolute top-0 left-0 w-full h-2 bg-gradient-to-r from-primary via-dark to-primary" />
        
        <div className="text-center mb-8">
           <h1 className="text-3xl font-black text-dark tracking-tighter">Welcome Back</h1>
           <p className="text-neutral font-medium mt-1">Sign in to your account</p>
        </div>

        {/* Role Selector */}
        <div className="flex bg-slate-100 p-1.5 rounded-full mb-8 relative">
           {['customer', 'vendor', 'admin'].map((r) => (
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
               left: role === 'customer' ? '0.375rem' : role === 'vendor' ? '33.33%' : '66.66%',
               width: 'calc(33.33% - 0.75rem)'
             }}
           />
        </div>

        <form onSubmit={handleLogin} className="space-y-4">
          <Input 
            type="email" 
            placeholder="Email Address" 
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
            icon={<Icon icon="solar:letter-bold-duotone" className="w-5 h-5 text-dark" />}
          />
          <Input 
            type="password" 
            placeholder="Password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
            icon={<Icon icon="solar:lock-password-bold-duotone" className="w-5 h-5 text-dark" />}
          />
          
          <div className="flex justify-end">
            <a href="#" className="text-xs font-bold text-primary hover:underline">Forgot password?</a>
          </div>

          <Button type="submit" className="w-full h-14" disabled={loading}>
            {loading ? <div className="w-6 h-6 border-2 border-white border-t-transparent rounded-full animate-spin" /> : 'Secure Sign In'}
          </Button>
        </form>

        <p className="text-center text-sm font-medium text-neutral mt-8 -mb-2">
           Don't have an account? <Link to="/register" className="text-dark font-black hover:underline">Create one</Link>
        </p>

      </div>
    </div>
  );
}
