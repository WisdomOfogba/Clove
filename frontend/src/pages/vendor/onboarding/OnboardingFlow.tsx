import { useOnboardingStore } from '../../../store';
import { motion } from 'motion/react';
import { Step1 } from './Step1';
import { Step2 } from './Step2';
import { Step3 } from './Step3';
import { Step4 } from './Step4';
import { Step5 } from './Step5';
import { Step6 } from './Step6';
import { Icon } from '@iconify/react';
import { cn } from '../../../lib/utils';

export function VendorOnboarding() {
  const { data } = useOnboardingStore();
  const currentStep = data.step;

  const steps = [
    { number: 1, title: 'Business' },
    { number: 2, title: 'Legal' },
    { number: 3, title: 'Proof' },
    { number: 4, title: 'Review' },
    { number: 5, title: 'Process' },
    { number: 6, title: 'Result' },
  ];

  const renderStep = () => {
    switch (currentStep) {
      case 1: return <Step1 />;
      case 2: return <Step2 />;
      case 3: return <Step3 />;
      case 4: return <Step4 />;
      case 5: return <Step5 />;
      case 6: return <Step6 />;
      default: return <Step1 />;
    }
  };

  return (
    <div className="max-w-3xl mx-auto space-y-8">
      <div className="space-y-2 text-center">
        <h1 className="text-3xl font-bold tracking-tight text-dark">Partner with CloveDelight</h1>
        <p className="text-neutral">Complete the "Proof of Life" verification to start selling.</p>
      </div>

      {/* Modern Stepper */}
      <div className="relative flex justify-between items-center px-4 overflow-x-auto pb-4 scrollbar-hide">
        {steps.map((step, idx) => (
          <div key={step.number} className="flex items-center gap-2 group flex-shrink-0">
            <div className={cn(
              "relative flex flex-col items-center gap-2 transition-all",
              currentStep === step.number ? "opacity-100" : "opacity-40"
            )}>
              <div className={cn(
                "w-10 h-10 rounded-xl flex items-center justify-center font-bold text-sm transition-all shadow-sm",
                currentStep > step.number 
                  ? "bg-primary text-white scale-90" 
                  : currentStep === step.number 
                    ? "bg-primary text-white ring-4 ring-primary/20" 
                    : "bg-white border border-slate-200 text-slate-400"
              )}>
                {currentStep > step.number ? <Icon icon="lucide:check" className="w-5 h-5" /> : step.number}
              </div>
              <span className={cn(
                "text-[10px] font-bold uppercase tracking-widest",
                currentStep === step.number ? "text-primary" : "text-slate-400"
              )}>
                {step.title}
              </span>
            </div>
            {idx < steps.length - 1 && (
              <div className={cn(
                "w-8 sm:w-16 h-px mx-1 bg-gradient-to-r",
                currentStep > step.number ? "from-primary to-primary" : "from-slate-200 to-slate-100"
              )} />
            )}
          </div>
        ))}
      </div>

      <div className="bg-white rounded-3xl border border-slate-100 shadow-xl shadow-slate-200/50 min-h-[500px]">
        {renderStep()}
      </div>
    </div>
  );
}
