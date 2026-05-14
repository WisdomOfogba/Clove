import { Link } from 'react-router-dom';
import { Icon } from '@iconify/react';

export function Footer() {
  return (
    <footer className="bg-dark text-slate-300">
      <div className="container mx-auto px-6 sm:px-12 py-20 max-w-7xl">
        <div className="grid grid-cols-2 md:grid-cols-4 gap-12 lg:gap-20">
          <div className="col-span-2 md:col-span-1 space-y-6">
            <Link to="/" className="flex items-center gap-2 group inline-flex">
               <div className="w-10 h-10 rounded-full bg-white text-dark flex items-center justify-center transition-transform group-hover:scale-95">
                  <Icon icon="solar:shop-2-bold" className="w-5 h-5 text-primary" />
               </div>
               <span className="font-extrabold text-2xl tracking-tight text-white">
                 CLOVE
               </span>
            </Link>
            <p className="text-slate-400 text-sm leading-relaxed font-medium max-w-xs">
              Fast, verified, and secure food delivery. Your favorite meals, monitored and delivered with uncompromising hygiene.
            </p>
            <div className="flex gap-3">
               {[ "solar:camera-bold", "solar:play-bold", "solar:letter-bold" ].map(icon => (
                 <div key={icon} className="w-10 h-10 rounded-full bg-slate-800 flex items-center justify-center text-slate-400 hover:text-white hover:bg-primary transition-all cursor-pointer">
                   <Icon icon={icon} className="w-5 h-5" />
                 </div>
               ))}
            </div>
          </div>
          <div>
            <h4 className="font-bold text-white mb-6 text-lg tracking-tight">Discover</h4>
            <ul className="space-y-4 text-sm font-medium">
              <li><Link to="/restaurants" className="hover:text-primary transition-colors">Order Food</Link></li>
              <li><Link to="/restaurants" className="hover:text-primary transition-colors">Popular Places</Link></li>
              <li><Link to="/vendor/onboarding" className="hover:text-primary transition-colors">Join as a Partner</Link></li>
            </ul>
          </div>
          <div>
            <h4 className="font-bold text-white mb-6 text-lg tracking-tight">Support</h4>
            <ul className="space-y-4 text-sm font-medium">
              <li><Link to="/help" className="hover:text-primary transition-colors">Help Center</Link></li>
              <li><Link to="/safety" className="hover:text-primary transition-colors">Safety Standard</Link></li>
              <li><Link to="/terms" className="hover:text-primary transition-colors">Terms of Service</Link></li>
              <li><Link to="/privacy" className="hover:text-primary transition-colors">Privacy Policy</Link></li>
            </ul>
          </div>
          <div>
            <h4 className="font-bold text-white mb-6 text-lg tracking-tight">Get the App</h4>
            <div className="space-y-4">
              <button className="flex items-center gap-3 bg-white text-dark px-5 py-3 rounded-full w-full hover:bg-slate-200 transition-all shadow-lg shadow-white/5">
                <Icon icon="ic:baseline-apple" className="w-6 h-6" />
                <div className="text-left">
                  <p className="text-[10px] leading-none font-bold text-slate-500">Download on the</p>
                  <p className="text-sm font-extrabold leading-none mt-1">App Store</p>
                </div>
              </button>
              <button className="flex items-center gap-3 bg-slate-800 text-white px-5 py-3 rounded-full w-full hover:bg-slate-700 transition-all border border-slate-700">
                <Icon icon="logos:google-play-icon" className="w-6 h-6" />
                <div className="text-left">
                  <p className="text-[10px] leading-none font-bold text-slate-400">GET IT ON</p>
                  <p className="text-sm font-extrabold leading-none mt-1">Google Play</p>
                </div>
              </button>
            </div>
          </div>
        </div>
        <div className="mt-16 pt-8 border-t border-slate-800 flex flex-col md:flex-row justify-between items-center gap-4">
          <p className="text-sm font-medium text-slate-500">© 2026 Clove App. All rights reserved.</p>
          <div className="flex items-center gap-6 text-sm font-medium text-slate-500">
             <Link to="/privacy" className="hover:text-white transition-colors">Privacy</Link>
             <Link to="/terms" className="hover:text-white transition-colors">Terms</Link>
          </div>
        </div>
      </div>
    </footer>
  );
}
