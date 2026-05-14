import { Icon } from '@iconify/react';
import { Link, useLocation } from 'react-router-dom';
import { useCart } from '../../contexts/CartContext';
import { motion } from 'motion/react';

export function Navbar() {
  const { totalItems } = useCart();
  const location = useLocation();

  const links = [
    { name: "Home", path: "/" },
    { name: "Restaurants", path: "/restaurants" },
    { name: "Partner With Us", path: "/vendor/onboarding" },
  ];

  return (
    <nav className="sticky top-0 z-50 bg-white/95 backdrop-blur-xl border-b border-slate-100">
      <div className="container mx-auto px-6 sm:px-12 h-24 flex items-center justify-between max-w-7xl">
        {/* Logo */}
        <Link to="/" className="flex items-center gap-2 group">
          <div className="w-10 h-10 rounded-full bg-dark text-white flex items-center justify-center transition-transform group-active:scale-95">
             <Icon icon="solar:shop-2-bold" className="w-5 h-5 text-primary" />
          </div>
          <span className="font-extrabold text-2xl tracking-tight text-dark">
            CLOVE
          </span>
        </Link>

        {/* Desktop Navigation */}
        <div className="hidden lg:flex items-center gap-8">
           {links.map(link => {
             const isActive = location.pathname === link.path;
             return (
               <Link 
                 key={link.name} 
                 to={link.path} 
                 className={`text-sm font-semibold transition-colors ${isActive ? 'text-primary' : 'text-slate-500 hover:text-dark'}`}
               >
                 {link.name}
               </Link>
             );
           })}
        </div>

        {/* Actions */}
        <div className="flex items-center gap-4">
          <Link to="/login" className="hidden sm:flex items-center justify-center font-bold text-sm text-dark hover:text-primary transition-colors">
            Sign In
          </Link>

          <Link to="/register" className="flex items-center gap-2 px-6 py-3.5 bg-primary text-white rounded-full font-bold text-sm hover:opacity-90 transition-all active:scale-95 shadow-[0_8px_20px_rgba(245,158,11,0.25)]">
            Sign Up
          </Link>

          {/* Cart with Count badge */}
          <Link to="/cart" className="relative group ml-4">
            <div className="w-12 h-12 rounded-full border border-slate-100 bg-white flex items-center justify-center text-dark group-hover:border-primary group-hover:text-primary transition-all shadow-sm">
              <Icon icon="solar:bag-3-outline" className="w-6 h-6" />
            </div>
            {totalItems > 0 && (
              <motion.span 
                initial={{ scale: 0 }}
                animate={{ scale: 1 }}
                className="absolute -top-1 -right-1 w-5 h-5 bg-dark text-white text-[10px] font-bold rounded-full flex items-center justify-center border-2 border-white"
              >
                {totalItems}
              </motion.span>
            )}
          </Link>
        </div>
      </div>
    </nav>
  );
}
