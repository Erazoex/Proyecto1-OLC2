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
