// Code generated from Grammar.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Grammar

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by GrammarParser.
type GrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GrammarParser#init.
	VisitInit(ctx *InitContext) interface{}

	// Visit a parse tree produced by GrammarParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by GrammarParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#declstmtWithTypeAndExpr.
	VisitDeclstmtWithTypeAndExpr(ctx *DeclstmtWithTypeAndExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#declstmtWithExpr.
	VisitDeclstmtWithExpr(ctx *DeclstmtWithExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#declstmtWithType.
	VisitDeclstmtWithType(ctx *DeclstmtWithTypeContext) interface{}

	// Visit a parse tree produced by GrammarParser#asignstmt.
	VisitAsignstmt(ctx *AsignstmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#incstmt.
	VisitIncstmt(ctx *IncstmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#decstmt.
	VisitDecstmt(ctx *DecstmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#ifSimple.
	VisitIfSimple(ctx *IfSimpleContext) interface{}

	// Visit a parse tree produced by GrammarParser#ifWithElse.
	VisitIfWithElse(ctx *IfWithElseContext) interface{}

	// Visit a parse tree produced by GrammarParser#ifWithElseIf.
	VisitIfWithElseIf(ctx *IfWithElseIfContext) interface{}

	// Visit a parse tree produced by GrammarParser#switchstmt.
	VisitSwitchstmt(ctx *SwitchstmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#switchcase.
	VisitSwitchcase(ctx *SwitchcaseContext) interface{}

	// Visit a parse tree produced by GrammarParser#printlnstmt.
	VisitPrintlnstmt(ctx *PrintlnstmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#exprWithParams.
	VisitExprWithParams(ctx *ExprWithParamsContext) interface{}

	// Visit a parse tree produced by GrammarParser#exprWithParam.
	VisitExprWithParam(ctx *ExprWithParamContext) interface{}

	// Visit a parse tree produced by GrammarParser#whilestmt.
	VisitWhilestmt(ctx *WhilestmtContext) interface{}

	// Visit a parse tree produced by GrammarParser#forWithExpr.
	VisitForWithExpr(ctx *ForWithExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#forWithRange.
	VisitForWithRange(ctx *ForWithRangeContext) interface{}

	// Visit a parse tree produced by GrammarParser#forrange.
	VisitForrange(ctx *ForrangeContext) interface{}

	// Visit a parse tree produced by GrammarParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by GrammarParser#array_def.
	VisitArray_def(ctx *Array_defContext) interface{}

	// Visit a parse tree produced by GrammarParser#primitiveType.
	VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{}

	// Visit a parse tree produced by GrammarParser#DoubleExpr.
	VisitDoubleExpr(ctx *DoubleExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#BoolExpr.
	VisitBoolExpr(ctx *BoolExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#IdExpr.
	VisitIdExpr(ctx *IdExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#ParExpr.
	VisitParExpr(ctx *ParExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#StrExpr.
	VisitStrExpr(ctx *StrExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#NotExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#IntExpr.
	VisitIntExpr(ctx *IntExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#nilExpr.
	VisitNilExpr(ctx *NilExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#OpExpr.
	VisitOpExpr(ctx *OpExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#UnaryExpr.
	VisitUnaryExpr(ctx *UnaryExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#CharExpr.
	VisitCharExpr(ctx *CharExprContext) interface{}
}
