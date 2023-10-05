// Code generated from SwiftGrammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // SwiftGrammar
import "github.com/antlr4-go/antlr/v4"

// SwiftGrammarListener is a complete listener for a parse tree produced by SwiftGrammarParser.
type SwiftGrammarListener interface {
	antlr.ParseTreeListener

	// EnterS is called when entering the s production.
	EnterS(c *SContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterGuardstatement is called when entering the guardstatement production.
	EnterGuardstatement(c *GuardstatementContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterStructfuncall is called when entering the structfuncall production.
	EnterStructfuncall(c *StructfuncallContext)

	// EnterStructinstruction is called when entering the structinstruction production.
	EnterStructinstruction(c *StructinstructionContext)

	// EnterStructfuncdec is called when entering the structfuncdec production.
	EnterStructfuncdec(c *StructfuncdecContext)

	// EnterFuncblock is called when entering the funcblock production.
	EnterFuncblock(c *FuncblockContext)

	// EnterFuncinst is called when entering the funcinst production.
	EnterFuncinst(c *FuncinstContext)

	// EnterSelfstructmodification is called when entering the selfstructmodification production.
	EnterSelfstructmodification(c *SelfstructmodificationContext)

	// EnterStructmodification is called when entering the structmodification production.
	EnterStructmodification(c *StructmodificationContext)

	// EnterStructblock is called when entering the structblock production.
	EnterStructblock(c *StructblockContext)

	// EnterStructdef is called when entering the structdef production.
	EnterStructdef(c *StructdefContext)

	// EnterRetturn is called when entering the retturn production.
	EnterRetturn(c *RetturnContext)

	// EnterFuncdec is called when entering the funcdec production.
	EnterFuncdec(c *FuncdecContext)

	// EnterVectormodification is called when entering the vectormodification production.
	EnterVectormodification(c *VectormodificationContext)

	// EnterForloop is called when entering the forloop production.
	EnterForloop(c *ForloopContext)

	// EnterRange is called when entering the range production.
	EnterRange(c *RangeContext)

	// EnterRemoveatvec is called when entering the removeatvec production.
	EnterRemoveatvec(c *RemoveatvecContext)

	// EnterAppendvec is called when entering the appendvec production.
	EnterAppendvec(c *AppendvecContext)

	// EnterRemovelastvec is called when entering the removelastvec production.
	EnterRemovelastvec(c *RemovelastvecContext)

	// EnterVecdec is called when entering the vecdec production.
	EnterVecdec(c *VecdecContext)

	// EnterBreakstatement is called when entering the breakstatement production.
	EnterBreakstatement(c *BreakstatementContext)

	// EnterContinuestatement is called when entering the continuestatement production.
	EnterContinuestatement(c *ContinuestatementContext)

	// EnterSwitchstatement is called when entering the switchstatement production.
	EnterSwitchstatement(c *SwitchstatementContext)

	// EnterCaselist is called when entering the caselist production.
	EnterCaselist(c *CaselistContext)

	// EnterCase is called when entering the case production.
	EnterCase(c *CaseContext)

	// EnterDefaultstatement is called when entering the defaultstatement production.
	EnterDefaultstatement(c *DefaultstatementContext)

	// EnterIfstmt is called when entering the ifstmt production.
	EnterIfstmt(c *IfstmtContext)

	// EnterEliflist is called when entering the eliflist production.
	EnterEliflist(c *EliflistContext)

	// EnterElif is called when entering the elif production.
	EnterElif(c *ElifContext)

	// EnterElsestament is called when entering the elsestament production.
	EnterElsestament(c *ElsestamentContext)

	// EnterPrintstmt is called when entering the printstmt production.
	EnterPrintstmt(c *PrintstmtContext)

	// EnterWhile_statement is called when entering the while_statement production.
	EnterWhile_statement(c *While_statementContext)

	// EnterVardec is called when entering the vardec production.
	EnterVardec(c *VardecContext)

	// EnterConstdec is called when entering the constdec production.
	EnterConstdec(c *ConstdecContext)

	// EnterAsignation is called when entering the asignation production.
	EnterAsignation(c *AsignationContext)

	// EnterUnarysum is called when entering the unarysum production.
	EnterUnarysum(c *UnarysumContext)

	// EnterUnarysub is called when entering the unarysub production.
	EnterUnarysub(c *UnarysubContext)

	// EnterIsemptyvec is called when entering the isemptyvec production.
	EnterIsemptyvec(c *IsemptyvecContext)

	// EnterCountvec is called when entering the countvec production.
	EnterCountvec(c *CountvecContext)

	// EnterVectoraccess is called when entering the vectoraccess production.
	EnterVectoraccess(c *VectoraccessContext)

	// EnterStructaccess is called when entering the structaccess production.
	EnterStructaccess(c *StructaccessContext)

	// EnterVecindexstruct is called when entering the vecindexstruct production.
	EnterVecindexstruct(c *VecindexstructContext)

	// EnterVectorstructaccess is called when entering the vectorstructaccess production.
	EnterVectorstructaccess(c *VectorstructaccessContext)

	// EnterIndexesList is called when entering the indexesList production.
	EnterIndexesList(c *IndexesListContext)

	// EnterVecac is called when entering the vecac production.
	EnterVecac(c *VecacContext)

	// EnterMatrix_type is called when entering the matrix_type production.
	EnterMatrix_type(c *Matrix_typeContext)

	// EnterRepeatingvector is called when entering the repeatingvector production.
	EnterRepeatingvector(c *RepeatingvectorContext)

	// EnterManualdef is called when entering the manualdef production.
	EnterManualdef(c *ManualdefContext)

	// EnterManualmatrixdef is called when entering the manualmatrixdef production.
	EnterManualmatrixdef(c *ManualmatrixdefContext)

	// EnterValues2 is called when entering the values2 production.
	EnterValues2(c *Values2Context)

	// EnterFuncarglist is called when entering the funcarglist production.
	EnterFuncarglist(c *FuncarglistContext)

	// EnterFuncarg is called when entering the funcarg production.
	EnterFuncarg(c *FuncargContext)

	// EnterFuncparameterlist is called when entering the funcparameterlist production.
	EnterFuncparameterlist(c *FuncparameterlistContext)

	// EnterFuncparameter is called when entering the funcparameter production.
	EnterFuncparameter(c *FuncparameterContext)

	// EnterDecmatrix is called when entering the decmatrix production.
	EnterDecmatrix(c *DecmatrixContext)

	// EnterAttrlist is called when entering the attrlist production.
	EnterAttrlist(c *AttrlistContext)

	// EnterAttr is called when entering the attr production.
	EnterAttr(c *AttrContext)

	// EnterStructexp is called when entering the structexp production.
	EnterStructexp(c *StructexpContext)

	// EnterKeyvaluelist is called when entering the keyvaluelist production.
	EnterKeyvaluelist(c *KeyvaluelistContext)

	// EnterKeyvalue is called when entering the keyvalue production.
	EnterKeyvalue(c *KeyvalueContext)

	// EnterCallfuncexp is called when entering the callfuncexp production.
	EnterCallfuncexp(c *CallfuncexpContext)

	// EnterCallfuncins is called when entering the callfuncins production.
	EnterCallfuncins(c *CallfuncinsContext)

	// EnterCast is called when entering the cast production.
	EnterCast(c *CastContext)

	// EnterSelfattributeaccess is called when entering the selfattributeaccess production.
	EnterSelfattributeaccess(c *SelfattributeaccessContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// ExitS is called when exiting the s production.
	ExitS(c *SContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitGuardstatement is called when exiting the guardstatement production.
	ExitGuardstatement(c *GuardstatementContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitStructfuncall is called when exiting the structfuncall production.
	ExitStructfuncall(c *StructfuncallContext)

	// ExitStructinstruction is called when exiting the structinstruction production.
	ExitStructinstruction(c *StructinstructionContext)

	// ExitStructfuncdec is called when exiting the structfuncdec production.
	ExitStructfuncdec(c *StructfuncdecContext)

	// ExitFuncblock is called when exiting the funcblock production.
	ExitFuncblock(c *FuncblockContext)

	// ExitFuncinst is called when exiting the funcinst production.
	ExitFuncinst(c *FuncinstContext)

	// ExitSelfstructmodification is called when exiting the selfstructmodification production.
	ExitSelfstructmodification(c *SelfstructmodificationContext)

	// ExitStructmodification is called when exiting the structmodification production.
	ExitStructmodification(c *StructmodificationContext)

	// ExitStructblock is called when exiting the structblock production.
	ExitStructblock(c *StructblockContext)

	// ExitStructdef is called when exiting the structdef production.
	ExitStructdef(c *StructdefContext)

	// ExitRetturn is called when exiting the retturn production.
	ExitRetturn(c *RetturnContext)

	// ExitFuncdec is called when exiting the funcdec production.
	ExitFuncdec(c *FuncdecContext)

	// ExitVectormodification is called when exiting the vectormodification production.
	ExitVectormodification(c *VectormodificationContext)

	// ExitForloop is called when exiting the forloop production.
	ExitForloop(c *ForloopContext)

	// ExitRange is called when exiting the range production.
	ExitRange(c *RangeContext)

	// ExitRemoveatvec is called when exiting the removeatvec production.
	ExitRemoveatvec(c *RemoveatvecContext)

	// ExitAppendvec is called when exiting the appendvec production.
	ExitAppendvec(c *AppendvecContext)

	// ExitRemovelastvec is called when exiting the removelastvec production.
	ExitRemovelastvec(c *RemovelastvecContext)

	// ExitVecdec is called when exiting the vecdec production.
	ExitVecdec(c *VecdecContext)

	// ExitBreakstatement is called when exiting the breakstatement production.
	ExitBreakstatement(c *BreakstatementContext)

	// ExitContinuestatement is called when exiting the continuestatement production.
	ExitContinuestatement(c *ContinuestatementContext)

	// ExitSwitchstatement is called when exiting the switchstatement production.
	ExitSwitchstatement(c *SwitchstatementContext)

	// ExitCaselist is called when exiting the caselist production.
	ExitCaselist(c *CaselistContext)

	// ExitCase is called when exiting the case production.
	ExitCase(c *CaseContext)

	// ExitDefaultstatement is called when exiting the defaultstatement production.
	ExitDefaultstatement(c *DefaultstatementContext)

	// ExitIfstmt is called when exiting the ifstmt production.
	ExitIfstmt(c *IfstmtContext)

	// ExitEliflist is called when exiting the eliflist production.
	ExitEliflist(c *EliflistContext)

	// ExitElif is called when exiting the elif production.
	ExitElif(c *ElifContext)

	// ExitElsestament is called when exiting the elsestament production.
	ExitElsestament(c *ElsestamentContext)

	// ExitPrintstmt is called when exiting the printstmt production.
	ExitPrintstmt(c *PrintstmtContext)

	// ExitWhile_statement is called when exiting the while_statement production.
	ExitWhile_statement(c *While_statementContext)

	// ExitVardec is called when exiting the vardec production.
	ExitVardec(c *VardecContext)

	// ExitConstdec is called when exiting the constdec production.
	ExitConstdec(c *ConstdecContext)

	// ExitAsignation is called when exiting the asignation production.
	ExitAsignation(c *AsignationContext)

	// ExitUnarysum is called when exiting the unarysum production.
	ExitUnarysum(c *UnarysumContext)

	// ExitUnarysub is called when exiting the unarysub production.
	ExitUnarysub(c *UnarysubContext)

	// ExitIsemptyvec is called when exiting the isemptyvec production.
	ExitIsemptyvec(c *IsemptyvecContext)

	// ExitCountvec is called when exiting the countvec production.
	ExitCountvec(c *CountvecContext)

	// ExitVectoraccess is called when exiting the vectoraccess production.
	ExitVectoraccess(c *VectoraccessContext)

	// ExitStructaccess is called when exiting the structaccess production.
	ExitStructaccess(c *StructaccessContext)

	// ExitVecindexstruct is called when exiting the vecindexstruct production.
	ExitVecindexstruct(c *VecindexstructContext)

	// ExitVectorstructaccess is called when exiting the vectorstructaccess production.
	ExitVectorstructaccess(c *VectorstructaccessContext)

	// ExitIndexesList is called when exiting the indexesList production.
	ExitIndexesList(c *IndexesListContext)

	// ExitVecac is called when exiting the vecac production.
	ExitVecac(c *VecacContext)

	// ExitMatrix_type is called when exiting the matrix_type production.
	ExitMatrix_type(c *Matrix_typeContext)

	// ExitRepeatingvector is called when exiting the repeatingvector production.
	ExitRepeatingvector(c *RepeatingvectorContext)

	// ExitManualdef is called when exiting the manualdef production.
	ExitManualdef(c *ManualdefContext)

	// ExitManualmatrixdef is called when exiting the manualmatrixdef production.
	ExitManualmatrixdef(c *ManualmatrixdefContext)

	// ExitValues2 is called when exiting the values2 production.
	ExitValues2(c *Values2Context)

	// ExitFuncarglist is called when exiting the funcarglist production.
	ExitFuncarglist(c *FuncarglistContext)

	// ExitFuncarg is called when exiting the funcarg production.
	ExitFuncarg(c *FuncargContext)

	// ExitFuncparameterlist is called when exiting the funcparameterlist production.
	ExitFuncparameterlist(c *FuncparameterlistContext)

	// ExitFuncparameter is called when exiting the funcparameter production.
	ExitFuncparameter(c *FuncparameterContext)

	// ExitDecmatrix is called when exiting the decmatrix production.
	ExitDecmatrix(c *DecmatrixContext)

	// ExitAttrlist is called when exiting the attrlist production.
	ExitAttrlist(c *AttrlistContext)

	// ExitAttr is called when exiting the attr production.
	ExitAttr(c *AttrContext)

	// ExitStructexp is called when exiting the structexp production.
	ExitStructexp(c *StructexpContext)

	// ExitKeyvaluelist is called when exiting the keyvaluelist production.
	ExitKeyvaluelist(c *KeyvaluelistContext)

	// ExitKeyvalue is called when exiting the keyvalue production.
	ExitKeyvalue(c *KeyvalueContext)

	// ExitCallfuncexp is called when exiting the callfuncexp production.
	ExitCallfuncexp(c *CallfuncexpContext)

	// ExitCallfuncins is called when exiting the callfuncins production.
	ExitCallfuncins(c *CallfuncinsContext)

	// ExitCast is called when exiting the cast production.
	ExitCast(c *CastContext)

	// ExitSelfattributeaccess is called when exiting the selfattributeaccess production.
	ExitSelfattributeaccess(c *SelfattributeaccessContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)
}
